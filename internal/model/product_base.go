package model

type BaseProduct struct {
	ID            uint32
	Description   string
	Count         uint32
	Title         string
	Price         uint32
	OriginalPrice uint32
	Discount      uint32
	Rating        float32
	Weight        float32
}
