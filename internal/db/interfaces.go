package db

type ProductManager interface {
	GetProductByID(id string) (Product, bool)
	GetAllProducts() map[string]Product
}

type UserManager interface {
	GetUserByEmail(email string) (User, bool)
	CreateUser(email string, user User) error
}
