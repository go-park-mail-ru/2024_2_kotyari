package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/db"
	"github.com/gorilla/mux"
)

type CardsApp struct {
	db db.ProductManager
}

func NewCardsApp(productsDB db.ProductManager) *CardsApp {
	return &CardsApp{
		db: productsDB,
	}
}

// Products
// @Summary Get Products
// @Description Возвращает список всех продуктов
// @Tags Products
// @Produce json
// @Success 200 {object} map[string]db.Product
// @Failure 500 {string} string "Ошибка при кодировании JSON"
// @Router /catalog/products [get]
func (c *CardsApp) Products(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products := c.db.GetAllProducts()

	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		http.Error(w, "Ошибка при кодировании JSON", http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

// ProductByID @Summary Get BaseProduct by ID
// @Description Возвращает продукт по его ID
// @Tags Products
// @Produce json
// @Param id path string true "BaseProduct ID"
// @Success 200 {object} db.Product
// @Failure 404 {string} string "Продукт не найден"
// @Failure 500 {string} string "Ошибка при кодировании JSON"
// @Router /catalog/product/{id} [get]
func (c *CardsApp) ProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	product, exists := c.db.GetProductByID(productID)
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
