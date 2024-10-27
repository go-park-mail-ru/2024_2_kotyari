package orders

import (
	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/google/uuid"
)

type getOrdersResponse struct {
	Status int        `json:"status"`
	Body   []orderDTO `json:"body"`
}

type getOrderByIDResponse struct {
	Status int      `json:"status"`
	Body   orderDTO `json:"body"`
}

type createOrderResponse struct {
	Status int      `json:"status"`
	Body   orderDTO `json:"body"`
}

type createOrderRequest struct {
	Address string `json:"address"`
}

type orderDTO struct {
	ID           uuid.UUID    `json:"id"`
	OrderDate    string       `json:"order_date"`
	DeliveryDate string       `json:"delivery_date"`
	Products     []productDTO `json:"products"`
	Address      string       `json:"address"`
}

type productDTO struct {
	ID       string `json:"id"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
	Cost     int    `json:"cost,omitempty"`
	Count    int    `json:"count,omitempty"`
	Weight   int    `json:"weight,omitempty"`
}

func ToOrderDTO(o order.Order) orderDTO {
	products := make([]productDTO, len(o.Products))
	for i, p := range o.Products {
		products[i] = productDTO{
			ID:       p.ID,
			ImageURL: p.ImageURL,
			Name:     p.Name,
			Cost:     p.Cost,
			Count:    p.Count,
			Weight:   p.Weight,
		}
	}

	return orderDTO{
		ID:           o.ID,
		OrderDate:    o.OrderDate.Format("2024-10-28"),
		DeliveryDate: o.DeliveryDate.Format("2024-10-28"),
		Products:     products,
	}
}

func convertOrdersToDTOs(orders []order.Order) []orderDTO {
	orderDTOs := make([]orderDTO, len(orders))
	for i, ord := range orders {
		orderDTOs[i] = ToOrderDTO(ord)
	}
	return orderDTOs
}
