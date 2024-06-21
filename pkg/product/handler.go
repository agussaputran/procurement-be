package product

import "context"

type IProdcutHandler interface {
	Store(ctx context.Context, request []ProductDataInput) (*ProductFetchResponse, error)
	Update(ctx context.Context, request []ProductDataInput) (*ProductFetchResponse, error)
	Fetch(ctx context.Context, request ProductDataInput) (*ProductFetchResponse, error)
}

type ProdcutHandler struct {
	usecase IProductUsecase
}

func NewProdcutHandler(usecase IProductUsecase) IProdcutHandler {
	return &ProdcutHandler{
		usecase: usecase,
	}
}

func (h *ProdcutHandler) Store(ctx context.Context, request []ProductDataInput) (response *ProductFetchResponse, err error) {
	response, err = h.usecase.Store(ctx, request)
	return
}

func (h *ProdcutHandler) Update(ctx context.Context, request []ProductDataInput) (response *ProductFetchResponse, err error) {
	response, err = h.usecase.Update(ctx, request)
	return
}

func (h *ProdcutHandler) Fetch(ctx context.Context, request ProductDataInput) (response *ProductFetchResponse, err error) {
	response, err = h.usecase.Fetch(ctx, request)
	return
}
