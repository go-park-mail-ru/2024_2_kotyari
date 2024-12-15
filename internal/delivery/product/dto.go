package product

import "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"

type dtoProductCatalog struct {
	dtoProduct
	ImageURL string `json:"image_url"`
}

type dtoProduct struct {
	ID            uint32  `json:"id"`
	Description   string  `json:"description"`
	Count         uint32  `json:"count"`
	Title         string  `json:"title"`
	Price         uint32  `json:"price"`
	OriginalPrice uint32  `json:"original_price"`
	Discount      uint32  `json:"discount"`
	Rating        float32 `json:"rating"`
}

func newDTOProductCatalogFromModel(pc model.ProductCatalog) dtoProductCatalog {
	return dtoProductCatalog{
		dtoProduct: dtoProduct{
			ID:            pc.ID,
			Description:   pc.Description,
			Count:         pc.Count,
			Title:         pc.Title,
			Price:         pc.Price,
			OriginalPrice: pc.OriginalPrice,
			Discount:      pc.Discount,
			Rating:        pc.Rating,
		},
		ImageURL: pc.ImageURL,
	}
}

type dtoImage struct {
	URL string `json:"url"`
}

type dtoOption struct {
	Link  string `json:"link"`
	Value string `json:"value"`
}

type dtoOptionsBlock struct {
	Title   string      `json:"title"`
	Type    string      `json:"type"`
	Options []dtoOption `json:"options"`
}

type dtoOptions struct {
	Values []dtoOptionsBlock `json:"values"`
}

type dtoSeller struct {
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type dtoProductCard struct {
	dtoProduct
	ReviewCount     uint32            `json:"review_count"`
	Images          []dtoImage        `json:"images"`
	Options         dtoOptions        `json:"options"`
	Characteristics map[string]string `json:"characteristics"`
	Seller          dtoSeller         `json:"seller"`
	InCart          bool              `json:"in_cart"`
}

func newDTOProductCardFromModel(pc model.ProductCard) dtoProductCard {
	images := make([]dtoImage, len(pc.Images))
	for i, img := range pc.Images {
		images[i] = dtoImage{URL: img.Url}
	}

	optionsBlocks := make([]dtoOptionsBlock, len(pc.Options.Values))
	for i, block := range pc.Options.Values {
		options := make([]dtoOption, len(block.Options))
		for j, opt := range block.Options {
			options[j] = dtoOption{
				Link:  opt.Link,
				Value: opt.Value,
			}
		}
		optionsBlocks[i] = dtoOptionsBlock{
			Title:   block.Title,
			Type:    block.Type,
			Options: options,
		}
	}

	return dtoProductCard{
		dtoProduct: dtoProduct{
			ID:            pc.ID,
			Description:   pc.Description,
			Count:         pc.Count,
			Title:         pc.Title,
			Price:         pc.Price,
			OriginalPrice: pc.OriginalPrice,
			Discount:      pc.Discount,
			Rating:        pc.Rating,
		},
		ReviewCount:     pc.ReviewCount,
		Images:          images,
		Options:         dtoOptions{Values: optionsBlocks},
		Characteristics: pc.Characteristics,
		Seller: dtoSeller{
			Name: pc.Seller.Name,
			Logo: pc.Seller.Logo,
		},
		InCart: pc.IsInCart,
	}
}
