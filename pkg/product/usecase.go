package product

import (
	"context"
)

type IProductUsecase interface {
	Store(ctx context.Context, arg string) error
}

type ProductUsecase struct {
	repository IProductRepository
}

func NewProductUsecase(repository IProductRepository) IProductUsecase {
	return &ProductUsecase{
		repository: repository,
	}
}

func (u *ProductUsecase) Store(ctx context.Context, arg string) (err error) {
	return
}
