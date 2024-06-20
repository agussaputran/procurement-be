package product

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type IProductRepository interface {
	Store(ctx context.Context, arg string) error
}

type ProductRepository struct {
	db *mongo.Client
}

func NewProductRepository(db *mongo.Client) IProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Store(ctx context.Context, arg string) (err error) {
	return
}
