package product

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

const queryGetImagesProduct = `
    SELECT 
        image_url
    FROM product_images
    WHERE product_id = $1;
`

func (ps *ProductsStore) getProductImages(ctx context.Context, productID uint64) ([]model.Image, error) {
	rows, err := ps.db.Query(ctx, queryGetImagesProduct, productID)
	if err != nil {
		ps.log.Error("[ ProductsStore.getProductImages ] ошибка выполнения запроса", "error", err.Error())
		return nil, err
	}
	defer rows.Close()

	var images []model.Image

	for rows.Next() {
		var image model.Image
		if err = rows.Scan(&image.Url); err != nil {
			ps.log.Error("[ ProductsStore.getProductImages ] ошибка чтения", "error", err.Error())
			return nil, err
		}

		images = append(images, image)
	}

	if len(images) == 0 {
		return nil, errs.ImagesDoesNotExists
	}

	return images, nil
}
