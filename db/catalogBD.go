package db

import "sync"

// Product представляет товар в каталоге
type Product struct {
	CurrentPrice     string `json:"currentPrice"`
	OldPrice         string `json:"oldPrice,omitempty"` // Поле не обязательно
	Discount         string `json:"discount,omitempty"` // Поле не обязательно
	Description      string `json:"description"`
	ShortDescription string `json:"shortDescription"`
	Image            string `json:"image"`
	URL              string `json:"url"`
	Currency         string `json:"currency,omitempty"` // Поле не обязательно
}

// ProductDB хранит продукты с поддержкой безопасного доступа через Mutex
type ProductDB struct {
	mu       sync.Mutex
	products map[string]Product
}

// GetProductByID возвращает продукт по ID
func GetProductByID(id string) (Product, bool) {
	productDB.mu.Lock()
	defer productDB.mu.Unlock()

	product, exists := productDB.products[id]
	return product, exists
}

// GetAllProducts возвращает все продукты
func GetAllProducts() map[string]Product {
	productDB.mu.Lock()
	defer productDB.mu.Unlock()

	productsCopy := make(map[string]Product)
	for id, product := range productDB.products {
		productsCopy[id] = product
	}
	return productsCopy
}
