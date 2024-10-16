package db

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
