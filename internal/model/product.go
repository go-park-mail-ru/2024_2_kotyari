package model

type Product struct {
	BaseProduct
	Categories []Category
	Type       string
	Tags       []string
}
