package model

type ProductInput struct {
	Name            string
	Description     string
	Count           uint64
	Price           uint32
	OriginalPrice   uint32
	Discount        uint32
	ImageURL        string
	Images          []string
	Characteristics map[string]string
	Options         *OptionsBlock
	CategoryIDs     []uint64
	Active          bool
}
