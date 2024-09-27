package handlers

import (
	"encoding/json"
	"net/http"

	"2024_2_kotyari/db"

	"github.com/gorilla/mux"
)

// @Summary Get Products
// @Description Возвращает список всех продуктов
// @Tags Products
// @Produce json
// @Success 200 {object} map[string]db.Product
// @Failure 500 {string} string "Ошибка при кодировании JSON"
// @Router /catalog/products [get]
func Products(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products := db.GetAllProducts()

	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		http.Error(w, "Ошибка при кодировании JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary Get Product by ID
// @Description Возвращает продукт по его ID
// @Tags Products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} db.Product
// @Failure 404 {string} string "Продукт не найден"
// @Failure 500 {string} string "Ошибка при кодировании JSON"
// @Router /catalog/product/{id} [get]
func ProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	product, exists := db.GetProductByID(productID)
	if !exists {
		http.Error(w, "Продукт не найден", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(product)
	if err != nil {
		http.Error(w, "Ошибка при кодировании JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
