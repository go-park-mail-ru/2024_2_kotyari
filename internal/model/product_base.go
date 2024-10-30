package model

type BaseProduct struct {
	ID            uint32
	Description   string
	Count         uint32
	Name          string
	Price         uint32
	OriginalPrice uint32
	Discount      uint32
	Rating        float32
}
