package role

import (
	"context"
)

type IRoleHandler interface {
	Store(ctx context.Context, request []RoleData) (*RoleFetchResponse, error)
	Update(ctx context.Context, request []RoleData) (*RoleFetchResponse, error)
	Fetch(ctx context.Context, request RoleData) (*RoleFetchResponse, error)
}

type RoleHandler struct {
	usecase IRoleUsecase
}

func NewRoleHandler(usecase IRoleUsecase) IRoleHandler {
	return &RoleHandler{
		usecase: usecase,
	}
}

func (h *RoleHandler) Store(ctx context.Context, request []RoleData) (response *RoleFetchResponse, err error) {
	response, err = h.usecase.Store(ctx, request)
	return
}

func (h *RoleHandler) Update(ctx context.Context, request []RoleData) (response *RoleFetchResponse, err error) {
	response, err = h.usecase.Update(ctx, request)
	return
}

func (h *RoleHandler) Fetch(ctx context.Context, request RoleData) (response *RoleFetchResponse, err error) {
	response, err = h.usecase.Fetch(ctx, request)
	return
}
