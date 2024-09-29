package db

import (
	"reflect"
	"testing"
)

func TestGetProductByID(t *testing.T) {
	db := NewProducts()

	t.Run("existing product", func(t *testing.T) {
		t.Parallel() // Параллельное выполнение теста

		// Проверяем существующий продукт
		id := "1"

		expectedProduct := productsData[id]

		product, exists := db.GetProductByID(id)
		if !exists {
			t.Errorf("expected product with ID %s to exist", id)
		}
		if !reflect.DeepEqual(product, expectedProduct) {
			t.Errorf("expected product %v, but got %v", expectedProduct, product)
		}
	})

	t.Run("non-existing product", func(t *testing.T) {
		t.Parallel() // Параллельное выполнение теста

		// Проверяем несуществующий продукт
		id := "999"
		_, exists := db.GetProductByID(id)
		if exists {
			t.Errorf("expected product with ID %s to not exist", id)
		}
	})
}

func TestGetAllProducts(t *testing.T) {
	db := NewProducts()
	t.Parallel() // Параллельное выполнение теста

	// Получаем все продукты
	products := db.GetAllProducts()

	// Проверяем количество продуктов
	expectedCount := len(db.products)
	if len(products) != expectedCount {
		t.Errorf("expected %d products, but got %d", expectedCount, len(products))
	}

	// Проверяем, что продукты совпадают
	for id, expectedProduct := range productsData {
		product, exists := products[id]
		if !exists {
			t.Errorf("expected product with ID %s to exist", id)
		}
		if !reflect.DeepEqual(product, expectedProduct) {
			t.Errorf("expected product %v, but got %v", expectedProduct, product)
		}
	}
}
