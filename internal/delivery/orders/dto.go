package orders

import (
	"github.com/google/uuid"
	"time"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type CreateOrderRequest struct {
	Address string `json:"address"`
}

type OrdersResponse struct {
	Orders []OrderResponse `json:"orders"`
}

type OrderResponse struct {
	ID           uuid.UUID    `json:"id"`
	OrderDate    time.Time    `json:"order_date"`
	DeliveryDate time.Time    `json:"delivery_date"`
	TotalPrice   uint16       `json:"total_price,omitempty"`
	Address      string       `json:"address"`
	Status       string       `json:"status,omitempty"`
	Products     []ProductDTO `json:"products"`
}

type OrderMaxResponse struct {
	ID           uuid.UUID    `json:"id"`
	OrderDate    time.Time    `json:"order_date"`
	DeliveryDate time.Time    `json:"delivery_date"`
	TotalPrice   uint16       `json:"total_price,omitempty"`
	Address      string       `json:"address"`
	Status       string       `json:"status,omitempty"`
	Recipient    string       `json:"recipient"`
	Products     []ProductDTO `json:"products"`
}

type ProductDTO struct {
	ID       uint32  `json:"id"`
	Cost     uint16  `json:"cost,omitempty"`
	Count    uint16  `json:"count,omitempty"`
	Weight   float32 `json:"weight,omitempty"`
	ImageURL string  `json:"image_url"`
	Name     string  `json:"name"`
}

func ToOrderResponse(o *order.Order) OrderResponse {
	products := make([]ProductDTO, 0, len(o.Products))

	for _, p := range o.Products {
		products = append(products, ProductDTO{
			ID:       p.ProductID,
			ImageURL: p.ImageUrl,
			Name:     p.Name,
			Cost:     p.Cost,
			Count:    p.Count,
			Weight:   p.Weight,
		})
	}

	return OrderResponse{
		ID:           o.ID,
		OrderDate:    o.OrderDate,
		DeliveryDate: o.DeliveryDate,
		Products:     products,
		Address:      o.Address,
		TotalPrice:   o.TotalPrice,
		Status:       o.Status,
	}
}

func ConvertOrdersToResponse(orders []order.Order) OrdersResponse {
	dtoOrders := make([]OrderResponse, 0, len(orders))

	for _, ord := range orders {
		dtoOrders = append(dtoOrders, ToOrderResponse(&ord))
	}

	return OrdersResponse{Orders: dtoOrders}
}

func ConvertProductsToDTO(products []order.ProductOrder) []ProductDTO {
	dto := make([]ProductDTO, 0, len(products))

	for _, p := range products {
		dto = append(dto, ProductDTO{
			ID:       p.ProductID,
			Cost:     p.Cost,
			Count:    p.Count,
			ImageURL: p.ImageUrl,
			Weight:   p.Weight,
			Name:     p.Name,
		})
	}

	return dto
}
