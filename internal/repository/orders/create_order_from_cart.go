package rorders

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

const defaultStatus = "awaiting_payment"

func (r *OrdersRepo) CreateOrderFromCart(ctx context.Context, orderData *order.OrderFromCart) (*order.Order, error) {
	_, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			defer tx.Rollback(ctx)
		}
	}()

	// Prepare statements
	createOrderStmt, err := tx.Prepare(ctx, "create_order", `
		INSERT INTO orders (id, user_id, total_price, address, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING created_at;
	`)
	if err != nil {
		r.logger.Error("[OrdersRepo.CreateOrderFromCart] failed to prepare create order query", slog.String("error", err.Error()))
		return nil, err
	}

	insertProductStmt, err := tx.Prepare(ctx, "insert_product_order", `
		INSERT INTO product_orders (id, order_id, product_id, option_id, count, delivery_date)
		SELECT $1, $2, $3, $4, $5, $6
		WHERE $5 > 0;
	`)
	if err != nil {
		r.logger.Error("[OrdersRepo.CreateOrderFromCart] failed to prepare insert product query", slog.String("error", err.Error()))
		return nil, err
	}

	removeCartItemsStmt, err := tx.Prepare(ctx, "remove_cart_items", `
		UPDATE carts
		SET count = 0, is_deleted = true
		WHERE user_id = $1 AND is_selected = true AND is_deleted = false;
	`)
	if err != nil {
		r.logger.Error("[OrdersRepo.CreateOrderFromCart] failed to prepare remove cart items query", slog.String("error", err.Error()))
		return nil, err
	}

	// Создаем заказ
	var createdAt time.Time
	err = tx.QueryRow(ctx, createOrderStmt.Name, orderData.OrderID, orderData.UserID, orderData.TotalPrice, orderData.Address).Scan(&orderData.CreatedAt)
	if err != nil {
		r.logger.Error("[OrdersRepo.CreateOrderFromCart] failed to insert order", slog.String("error", err.Error()), slog.Uint64("user_id", uint64(orderData.UserID)))
		return nil, err
	}

	// Батчирование запросов для вставки продуктов в заказ
	batch := &pgx.Batch{}
	for _, p := range orderData.Products {
		productOrderID := uuid.New()
		batch.Queue(insertProductStmt.Name, productOrderID, orderData.OrderID, p.ID, p.OptionID, p.Count, orderData.DeliveryDate)
	}

	// Выполнение батча
	br := tx.SendBatch(ctx, batch)

	for range orderData.Products {
		_, err = br.Exec()
		if err != nil {
			r.logger.Error("[OrdersRepo.CreateOrderFromCart] failed to insert product in order",
				slog.String("error", err.Error()),
				slog.Uint64("user_id", uint64(orderData.UserID)))
			return nil, err
		}
	}

	if err := br.Close(); err != nil {
		r.logger.Error("[OrdersRepo.CreateOrderFromCart] failed to close batch result",
			slog.String("error", err.Error()),
			slog.Uint64("user_id", uint64(orderData.UserID)))
		return nil, err
	}

	// Удаление товаров из корзины
	_, err = tx.Exec(ctx, removeCartItemsStmt.Name, orderData.UserID)
	if err != nil {
		r.logger.Error("[OrdersRepo.CreateOrderFromCart] failed to remove selected cart items",
			slog.String("error", err.Error()),
			slog.Uint64("user_id", uint64(orderData.UserID)))
		return nil, err
	}

	// Завершаем транзакцию
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &order.Order{
		ID:         orderData.OrderID,
		Address:    orderData.Address,
		Status:     defaultStatus,
		OrderDate:  createdAt,
		Products:   orderData.Products,
		TotalPrice: orderData.TotalPrice,
	}, nil
}
