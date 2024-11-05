package product

import (
	"context"
	"log/slog"
	"os"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
)

type imagesUsecase interface {
	SaveImage(filename string, file *os.File) (string, error)
}

type cartChecker interface {
	ProductInCart(ctx context.Context, userId uint32, productId uint32) (bool, error)
}

type productsRepo interface {
	AddProduct(ctx context.Context, card model.ProductCard) error
	GetAllProducts(ctx context.Context) ([]model.ProductCatalog, error)
	GetProductByID(ctx context.Context, productID uint64) (model.ProductCard, error)
}

type ProductsDelivery struct {
	repo          productsRepo
	log           *slog.Logger
	checker       cartChecker
	errResolver   errs.GetErrorCode
	imagesUsecase imagesUsecase
}

func NewProductHandler(errResolver errs.GetErrorCode, repo productsRepo, log *slog.Logger, cartChecker cartChecker, imagesUsecase imagesUsecase) *ProductsDelivery {
	return &ProductsDelivery{
		repo:          repo,
		log:           log,
		checker:       cartChecker,
		errResolver:   errResolver,
		imagesUsecase: imagesUsecase,
	}
}
