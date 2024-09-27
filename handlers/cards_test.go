package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"2024_2_kotyari/handlers"

	"github.com/gorilla/mux"
)

// Тест для Products
func TestProducts(t *testing.T) {
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.Products)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ProductsHandler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("ProductsHandler returned wrong Content-Type: got %v want %v",
			contentType, expectedContentType)
	}
}

// Тест на наличие ProductByID с действительным идентификатором продукта
func TestProductByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/product/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/product/{id}", handlers.ProductByID)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ProductByIDHandler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("ProductByIDHandler returned wrong Content-Type: got %v want %v",
			contentType, expectedContentType)
	}
}

// Тест для ProductByID с несуществующим идентификатором продукта
func TestProductByID_NotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/product/999", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/product/{id}", handlers.ProductByID)
	router.ServeHTTP(rr, req)

	// Проверяет, что код состояния соответствует ожиданиям.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("ProductByIDHandler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
