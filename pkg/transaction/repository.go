package transaction

import (
	"context"
	"log"
	"procurement-be/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ITransactionRepository interface {
	Store(ctx context.Context, args bson.M) error
	Update(ctx context.Context, filter, args bson.M) error
	Fetch(ctx context.Context, filter bson.M) ([]*TransactionData, error)
}

type TransactionRepository struct {
	db *mongo.Client
}

func NewTransactionRepository(db *mongo.Client) ITransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

// Store... insert data to transactions collection based on arg
func (r *TransactionRepository) Store(ctx context.Context, args bson.M) (err error) {
	// Set context with timeout for the insert operation
	ctxDB, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := utils.GetCollection("transactions", r.db)
	collection.InsertOne(ctxDB, args)
	return
}

// Update... update data to transactions collection based on arg
func (r *TransactionRepository) Update(ctx context.Context, filter, args bson.M) (err error) {
	// Set context with timeout for the update operation
	ctxDB, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := utils.GetCollection("transactions", r.db)
	collection.UpdateOne(ctxDB, filter, bson.M{"$set": args})
	return
}

// Fetch... select data from transactions collection based on filter
func (r *TransactionRepository) Fetch(ctx context.Context, filter bson.M) (data []*TransactionData, err error) {
	collection := utils.GetCollection("transactions", r.db)
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user TransactionData
		if err = cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		data = append(data, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
