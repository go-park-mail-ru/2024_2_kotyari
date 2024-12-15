package model

type ProductCard struct {
	Product
	ReviewCount     uint32
	Images          []Image
	Options         Options
	Characteristics map[string]string
	Seller          Seller
	IsInCart        bool
}

type Image struct {
	Url string
}

type OptionsBlock struct {
	Title   string
	Type    string
	Options []Option
}

type Option struct {
	Link  string
	Value string
}

type Options struct {
	Values []OptionsBlock
}
