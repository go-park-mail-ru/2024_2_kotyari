package product

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

const queryGetImagesProduct = `
    SELECT 
        image_url
    FROM product_images
    WHERE product_id = $1;
`

func (ps *ProductsStore) getProductImages(ctx context.Context, productID uint64) ([]model.Image, error) {
	rowsImages, err := ps.db.Query(ctx, queryGetImagesProduct, productID)
	if err != nil {
		ps.log.Error("[ ProductsStore.GetProductByID ] Error executing images query", "error", err.Error())
		return nil, err
	}
	defer rowsImages.Close()

	var images []model.Image

	for rowsImages.Next() {
		var image model.Image
		if err = rowsImages.Scan(&image.Url); err != nil {
			ps.log.Error("[ ProductsStore.GetProductByID ] Error scanning image", "error", err.Error())
			return nil, err
		}
		images = append(images, image)
	}

	if rowsImages.Err() != nil {
		ps.log.Error("[ ProductsStore.GetProductByID ] Error iterating over images rows", "error", rowsImages.Err().Error())
		return nil, rowsImages.Err()
	}

	return images, nil
}
