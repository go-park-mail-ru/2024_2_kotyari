package model

type CartProductsForOrderWithUser struct {
	Items                  []CartProductForOrder
	UserName               string
	PreferredPaymentMethod string
}
