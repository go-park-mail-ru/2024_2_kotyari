package db

type ProductManager interface {
	GetProductByID(id string) (Product, bool)
	GetAllProducts() map[string]Product
}
