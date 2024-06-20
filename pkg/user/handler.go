package user

import (
	"context"
)

type IUserHandler interface {
	Store(ctx context.Context, request UserData) (*LoginResponse, error)
	Update(ctx context.Context, request UserData) (*LoginResponse, error)
	Fetch(ctx context.Context, request UserData) (*UserFetchResponse, error)
	Login(ctx context.Context, request LoginRequest) (*LoginResponse, error)
}

type UserHandler struct {
	usecase IUserUsecase
}

func NewUserHandler(usecase IUserUsecase) IUserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (h *UserHandler) Store(ctx context.Context, request UserData) (response *LoginResponse, err error) {
	response, err = h.usecase.Store(ctx, request)
	return
}

func (h *UserHandler) Update(ctx context.Context, request UserData) (response *LoginResponse, err error) {
	response, err = h.usecase.Update(ctx, request)
	return
}

func (h *UserHandler) Fetch(ctx context.Context, request UserData) (response *UserFetchResponse, err error) {
	response, err = h.usecase.Fetch(ctx, request)
	return
}

func (h *UserHandler) Login(ctx context.Context, request LoginRequest) (response *LoginResponse, err error) {
	response, err = h.usecase.Login(ctx, request)
	return
}
