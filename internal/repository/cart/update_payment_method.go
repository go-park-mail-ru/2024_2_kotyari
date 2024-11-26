package cart

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
)

func (cs *CartsStore) UpdatePaymentMethod(ctx context.Context, userID uint32, method string) error {
	requestID, err := utils.GetContextRequestID(ctx)
	if err != nil {
		return err
	}

	cs.log.Info("[CartsStore.UpdatePaymentMethod] Started executing", slog.Any("request-id", requestID))

	const query = `UPDATE users SET preferred_payment_method = $1 WHERE id = $2`
	_, err = cs.db.Exec(ctx, query, method, userID)
	if err != nil {
		cs.log.Error("[CartsStore.UpdateUserPaymentMethod] Error updating payment method")
		return err
	}
	return nil
}
