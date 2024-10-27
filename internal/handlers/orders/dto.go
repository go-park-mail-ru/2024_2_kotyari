package orders

import (
	"github.com/google/uuid"

	order "github.com/go-park-mail-ru/2024_2_kotyari/internal/model/orders"
)

type GetOrdersResponse struct {
	Status int        `json:"status"`
	Body   []OrderDTO `json:"body"`
}

type GetOrderByIDResponse struct {
	Status int      `json:"status"`
	Body   OrderDTO `json:"body"`
}

type OrderDTO struct {
	ID           uuid.UUID    `json:"id"`
	OrderDate    string       `json:"order_date"`
	DeliveryDate string       `json:"delivery_date"`
	Products     []ProductDTO `json:"products"`
}

type ProductDTO struct {
	ID       string `json:"id"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
	Cost     int    `json:"cost,omitempty"`
	Count    int    `json:"count,omitempty"`
	Weight   int    `json:"weight,omitempty"`
}

func ToOrderDTO(o order.Order) OrderDTO {
	products := make([]ProductDTO, len(o.Products))
	for i, p := range o.Products {
		products[i] = ProductDTO{
			ID:       p.ID,
			ImageURL: p.ImageURL,
			Name:     p.Name,
			Cost:     p.Cost,
			Count:    p.Count,
			Weight:   p.Weight,
		}
	}

	return OrderDTO{
		ID:           o.ID,
		OrderDate:    o.OrderDate.Format("2024-10-28"),
		DeliveryDate: o.DeliveryDate.Format("2024-10-28"),
		Products:     products,
	}
}
