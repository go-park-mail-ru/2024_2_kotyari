package csrf

import (
	"os"
)

type CscfUsecase struct {
	secret string `env:"CSRF_SECRET,required"`
}

func NewCscfUsecase() *CscfUsecase {
	secret := os.Getenv("CSRF_SECRET")

	return &CscfUsecase{
		secret: secret,
	}
}
