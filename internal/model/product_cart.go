package model

type CartProduct struct {
	BaseProduct
	ImageURL   string
	IsSelected bool
	IsDeleted  bool
}
