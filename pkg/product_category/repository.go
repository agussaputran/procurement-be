package product_category

import (
	"context"
	"log"
	"procurement-be/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProductCategoryRepository interface {
	Fetch(ctx context.Context, filter bson.M) ([]*ProductCategoryData, error)
	Store(ctx context.Context, args []interface{}) error
	Update(ctx context.Context, args []mongo.WriteModel) error
}

type ProductCategoryRepository struct {
	db *mongo.Client
}

func NewProductCategoryRepository(db *mongo.Client) IProductCategoryRepository {
	return &ProductCategoryRepository{
		db: db,
	}
}

// Fetch... select data from roles collection based on filter
func (r *ProductCategoryRepository) Fetch(ctx context.Context, filter bson.M) (data []*ProductCategoryData, err error) {
	collection := utils.GetCollection("product_categories", r.db)
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var role ProductCategoryData
		if err = cursor.Decode(&role); err != nil {
			log.Fatal(err)
		}
		data = append(data, &role)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

// Store... insert data to roles collection based on arg
func (r *ProductCategoryRepository) Store(ctx context.Context, args []interface{}) (err error) {
	// Set context with timeout for the insert operation
	ctxDB, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := utils.GetCollection("product_categories", r.db)
	collection.InsertMany(ctxDB, args)
	return
}

// Update... update data to roles collection based on arg
func (r *ProductCategoryRepository) Update(ctx context.Context, args []mongo.WriteModel) (err error) {
	// Set context with timeout for the update operation
	ctxDB, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := utils.GetCollection("product_categories", r.db)
	collection.BulkWrite(ctxDB, args)
	return
}
