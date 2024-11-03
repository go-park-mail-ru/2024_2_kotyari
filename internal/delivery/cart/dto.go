package cart

import (
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type ChangeCartProductCountRequest struct {
	Count int32 `json:"count"`
}

func (r ChangeCartProductCountRequest) ToModel() int32 {
	return r.Count
}

type ChangeCartProductSelectedStateRequest struct {
	IsSelected bool `json:"is_selected"`
}

func (r ChangeCartProductSelectedStateRequest) ToModel() bool {
	return r.IsSelected
}

type GetCartResponse struct {
	ID           uint32               `json:"id"`
	DeliveryDate time.Time            `json:"delivery_date"`
	Products     []GetProductResponse `json:"products"`
}

type GetProductResponse struct {
	ID            uint32  `json:"id"`
	Description   string  `json:"description"`
	Count         uint32  `json:"count"`
	ImageURL      string  `json:"image_url"`
	IsSelected    bool    `json:"is_selected"`
	Title         string  `json:"title"`
	Price         uint32  `json:"price"`
	OriginalPrice uint32  `json:"original_price"`
	Discount      uint32  `json:"discount"`
	Rating        float32 `json:"rating"`
}

func cartResponseFromModel(cart model.Cart, products []GetProductResponse) GetCartResponse {
	return GetCartResponse{
		ID:           cart.ID,
		DeliveryDate: cart.DeliveryDate,
		Products:     products,
	}
}

func productResponseFromModel(cartProduct model.CartProduct) GetProductResponse {
	return GetProductResponse{
		ID:            cartProduct.ID,
		Description:   cartProduct.Description,
		Count:         cartProduct.Count,
		ImageURL:      cartProduct.ImageURL,
		IsSelected:    cartProduct.IsSelected,
		Title:         cartProduct.Title,
		Price:         cartProduct.Price,
		OriginalPrice: cartProduct.OriginalPrice,
		Discount:      cartProduct.Discount,
		Rating:        cartProduct.Rating,
	}
}
