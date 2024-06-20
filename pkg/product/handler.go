package product

import "context"

type IProdcutHandler interface {
	Store(ctx context.Context, arg string) error
}

type ProdcutHandler struct {
	usecase IProductUsecase
}

func NewProdcutHandler(usecase IProductUsecase) IProdcutHandler {
	return &ProdcutHandler{
		usecase: usecase,
	}
}

func (h *ProdcutHandler) Store(ctx context.Context, arg string) (err error) {
	return
}
