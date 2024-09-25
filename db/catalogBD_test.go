package db

import (
	"reflect"
	"testing"
)

func TestGetProductByID(t *testing.T) {
	t.Run("existing product", func(t *testing.T) {
		t.Parallel() // Параллельное выполнение теста

		// Проверяем существующий продукт
		id := "1"
		expectedProduct := Product{
			CurrentPrice:     "1999",
			OldPrice:         "2999",
			Discount:         "33",
			Description:      "Быстрая зарядка Type C, блок питания.",
			ShortDescription: "Быстрая зарядка Type C, блок питания.",
			URL:              "/catalog/product/1",
			Currency:         "$",
		}

		product, exists := GetProductByID(id)
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
		_, exists := GetProductByID(id)
		if exists {
			t.Errorf("expected product with ID %s to not exist", id)
		}
	})
}

func TestGetAllProducts(t *testing.T) {
	t.Parallel() // Параллельное выполнение теста

	// Получаем все продукты
	products := GetAllProducts()

	// Проверяем количество продуктов
	expectedCount := len(productDB.products)
	if len(products) != expectedCount {
		t.Errorf("expected %d products, but got %d", expectedCount, len(products))
	}

	// Проверяем, что продукты совпадают
	for id, expectedProduct := range productDB.products {
		product, exists := products[id]
		if !exists {
			t.Errorf("expected product with ID %s to exist", id)
		}
		if !reflect.DeepEqual(product, expectedProduct) {
			t.Errorf("expected product %v, but got %v", expectedProduct, product)
		}
	}
}
