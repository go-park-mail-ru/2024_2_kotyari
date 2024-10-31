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

type CartResponse struct {
	ID           uint32            `json:"id"`
	DeliveryDate time.Time         `json:"delivery_date"`
	Products     []ProductResponse `json:"products"`
}

type ProductResponse struct {
	ID            uint32  `json:"id"`
	Description   string  `json:"description"`
	Count         uint32  `json:"count"`
	ImageURL      string  `json:"image_url"`
	Title         string  `json:"title"`
	Price         uint32  `json:"price"`
	OriginalPrice uint32  `json:"original_price"`
	Discount      uint32  `json:"discount"`
	Rating        float32 `json:"rating"`
}

func CartResponseFromModel(cart model.Cart, products []ProductResponse) CartResponse {
	return CartResponse{
		ID:           cart.ID,
		DeliveryDate: cart.DeliveryDate,
		Products:     products,
	}
}

func ProductResponseFromModel(cartProduct model.CartProduct) ProductResponse {
	return ProductResponse{
		ID:            cartProduct.ID,
		Description:   cartProduct.Description,
		Count:         cartProduct.Count,
		ImageURL:      cartProduct.ImageURL,
		Title:         cartProduct.Title,
		Price:         cartProduct.Price,
		OriginalPrice: cartProduct.OriginalPrice,
		Discount:      cartProduct.Discount,
		Rating:        cartProduct.Rating,
	}
}
