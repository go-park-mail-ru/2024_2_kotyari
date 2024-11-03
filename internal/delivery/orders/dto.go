package orders

import (
	"github.com/google/uuid"
	"time"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type createOrderRequest struct {
	Address string `json:"address"`
}

type orderDTOs struct {
	Orders []orderDTO `json:"orders"`
}

type orderDTO struct {
	ID           uuid.UUID    `json:"id"`
	OrderDate    time.Time    `json:"order_date"`
	DeliveryDate time.Time    `json:"delivery_date"`
	TotalPrice   uint16       `json:"total_price,omitempty"`
	Address      string       `json:"address"`
	Status       string       `json:"status,omitempty"`
	Products     []productDTO `json:"products"`
}

type orderDTOMax struct {
	ID           uuid.UUID    `json:"id"`
	OrderDate    time.Time    `json:"order_date"`
	DeliveryDate time.Time    `json:"delivery_date"`
	TotalPrice   uint16       `json:"total_price,omitempty"`
	Address      string       `json:"address"`
	Status       string       `json:"status,omitempty"`
	Recipient    string       `json:"recipient"`
	Products     []productDTO `json:"products"`
}

type productDTO struct {
	ID       uint32 `json:"id"`
	Cost     uint16 `json:"cost,omitempty"`
	Count    uint16 `json:"count,omitempty"`
	Weight   uint16 `json:"weight,omitempty"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
}

func ToOrderDTO(o *order.Order) orderDTO {
	products := make([]productDTO, len(o.Products))

	for i, p := range o.Products {
		products[i] = productDTO{
			ID:       p.ProductID,
			ImageURL: p.ImageUrl,
			Name:     p.Name,
			Cost:     p.Cost,
			Count:    p.Count,
			Weight:   p.Weight,
		}
	}

	return orderDTO{
		ID:           o.ID,
		OrderDate:    o.OrderDate,
		DeliveryDate: o.DeliveryDate,
		Products:     products,
		Address:      o.Address,
		TotalPrice:   o.TotalPrice,
		Status:       o.Status,
	}
}

func convertOrdersToDTOs(orders []order.Order) orderDTOs {
	dtoOrders := make([]orderDTO, len(orders))

	for i, ord := range orders {
		dtoOrders[i] = ToOrderDTO(&ord)
	}

	return orderDTOs{Orders: dtoOrders}
}

func convertProductsToDTO(products []order.ProductOrder) []productDTO {
	dto := make([]productDTO, len(products))

	for i, p := range products {
		dto[i] = productDTO{
			ID:       p.ProductID,
			Cost:     p.Cost,
			Count:    p.Count,
			ImageURL: p.ImageUrl,
			Weight:   p.Weight,
			Name:     p.Name,
		}
	}

	return dto
}
