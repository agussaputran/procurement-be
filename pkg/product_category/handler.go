package product_category

import (
	"context"
)

type IProductCategoryHandler interface {
	Store(ctx context.Context, request []ProductCategoryData) (*ProductCategoryFetchResponse, error)
	Update(ctx context.Context, request []ProductCategoryData) (*ProductCategoryFetchResponse, error)
	Fetch(ctx context.Context, request ProductCategoryData) (*ProductCategoryFetchResponse, error)
}

type ProductCategoryHandler struct {
	usecase IProductCategoryUsecase
}

func NewProductCategoryHandler(usecase IProductCategoryUsecase) IProductCategoryHandler {
	return &ProductCategoryHandler{
		usecase: usecase,
	}
}

func (h *ProductCategoryHandler) Store(ctx context.Context, request []ProductCategoryData) (response *ProductCategoryFetchResponse, err error) {
	response, err = h.usecase.Store(ctx, request)
	return
}

func (h *ProductCategoryHandler) Update(ctx context.Context, request []ProductCategoryData) (response *ProductCategoryFetchResponse, err error) {
	response, err = h.usecase.Update(ctx, request)
	return
}

func (h *ProductCategoryHandler) Fetch(ctx context.Context, request ProductCategoryData) (response *ProductCategoryFetchResponse, err error) {
	response, err = h.usecase.Fetch(ctx, request)
	return
}
