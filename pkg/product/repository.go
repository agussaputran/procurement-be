package product

import (
	"context"
	"log"
	"procurement-be/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProductRepository interface {
	Store(ctx context.Context, args []interface{}) error
	Update(ctx context.Context, args []mongo.WriteModel) error
	Fetch(ctx context.Context, filter bson.M) ([]*ProductData, error)
}

type ProductRepository struct {
	db *mongo.Client
}

func NewProductRepository(db *mongo.Client) IProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Store(ctx context.Context, args []interface{}) (err error) {
	// Set context with timeout for the insert operation
	ctxDB, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := utils.GetCollection("products", r.db)
	collection.InsertMany(ctxDB, args)
	return
}

// Update... update data to products collection based on arg
func (r *ProductRepository) Update(ctx context.Context, args []mongo.WriteModel) (err error) {
	// Set context with timeout for the update operation
	ctxDB, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := utils.GetCollection("products", r.db)
	collection.BulkWrite(ctxDB, args)
	return
}

func (r *ProductRepository) Fetch(ctx context.Context, filter bson.M) (data []*ProductData, err error) {
	collection := utils.GetCollection("products", r.db)

	// Define the aggregation pipeline
	pipeline := mongo.Pipeline{
		{{"$match", filter}}, // Add your filter condition here
		{{"$lookup", bson.D{
			{"from", "users"},
			{"localField", "user_id"},
			{"foreignField", "id"},
			{"as", "vendor"},
		}}},
		{{"$unwind", "$vendor"}},
		{{"$lookup", bson.D{
			{"from", "product_categories"},
			{"localField", "category_id"},
			{"foreignField", "id"},
			{"as", "category"},
		}}},
		{{"$unwind", "$category"}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product ProductData
		if err = cursor.Decode(&product); err != nil {
			log.Fatal(err)
		}
		data = append(data, &product)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
