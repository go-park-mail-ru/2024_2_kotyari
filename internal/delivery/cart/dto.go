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

type orderData struct {
	TotalItems     uint16             `json:"total_items"`
	TotalWeight    float32            `json:"total_weight"`
	FinalPrice     float32            `json:"final_price"`
	Currency       string             `json:"currency"`
	PaymentMethods []paymentMethod    `json:"payment_methods"`
	Recipient      recipientInfo      `json:"recipient"`
	DeliveryDates  []deliveryDateInfo `json:"delivery_dates"`
}

type paymentMethod struct {
	Method     string `json:"method"`
	Icon       string `json:"icon"`
	IsSelected bool   `json:"is_selected"`
}

type recipientInfo struct {
	Address       string `json:"address"`
	RecipientName string `json:"recipient_name"`
}

type deliveryDateInfo struct {
	Date   time.Time         `json:"date"`
	Weight float32           `json:"weight"`
	Items  []productResponse `json:"items"`
}

type productResponse struct {
	Title    string  `json:"product_name"`
	Price    float32 `json:"product_price"`
	Quantity uint16  `json:"quantity"`
	Image    string  `json:"product_image"`
	Weight   float32 `json:"weight"`
	URL      string  `json:"url"`
}

type requestPaymentMethod struct {
	PaymentMethod string `json:"payment_method"`
}
