package db

import "sync"

// Products хранит продукты с поддержкой безопасного доступа через Mutex
type Products struct {
	mu       sync.RWMutex
	products map[string]Product
}

func NewProducts() *Products {
	return &Products{
		products: productsData,
	}
}

// GetProductByID возвращает продукт по ID
func (p *Products) GetProductByID(id string) (Product, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	product, exists := p.products[id]

	return product, exists
}

// GetAllProducts возвращает все продукты
func (p *Products) GetAllProducts() map[string]Product {
	p.mu.Lock()
	defer p.mu.Unlock()

	productsCopy := make(map[string]Product, len(p.products))
	for id, product := range p.products {
		productsCopy[id] = product
	}

	return productsCopy
}
