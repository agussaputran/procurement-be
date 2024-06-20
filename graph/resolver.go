package graph

import (
	"procurement-be/pkg/product"
	"procurement-be/pkg/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PkgHandler PkgHandler
}

type PkgHandler struct {
	UserHandler    user.IUserHandler
	ProductHandler product.IProdcutHandler
}
