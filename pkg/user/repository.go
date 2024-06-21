package user

import (
	"context"
	"log"
	"procurement-be/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	Store(ctx context.Context, args bson.M) error
	Update(ctx context.Context, filter, args bson.M) error
	Fetch(ctx context.Context, filter bson.M) ([]*UserData, error)
	FetchByEmail(ctx context.Context, email string) (*UserData, error)
}

type UserRepository struct {
	db *mongo.Client
}

func NewUserRepository(db *mongo.Client) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

// Store... insert data to users collection based on arg
func (r *UserRepository) Store(ctx context.Context, args bson.M) (err error) {
	// Set context with timeout for the insert operation
	ctxDB, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := utils.GetCollection("users", r.db)
	collection.InsertOne(ctxDB, args)
	return
}

// Update... update data to users collection based on arg
func (r *UserRepository) Update(ctx context.Context, filter, args bson.M) (err error) {
	// Set context with timeout for the update operation
	ctxDB, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := utils.GetCollection("users", r.db)
	collection.UpdateOne(ctxDB, filter, bson.M{"$set": args})
	return
}

// Fetch... select data from users collection based on filter
func (r *UserRepository) Fetch(ctx context.Context, filter bson.M) (data []*UserData, err error) {
	collection := utils.GetCollection("users", r.db)
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user UserData
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

func (r *UserRepository) FetchByEmail(ctx context.Context, email string) (data *UserData, err error) {
	collection := utils.GetCollection("users", r.db)
	cursor := collection.FindOne(ctx, bson.M{"email": email})

	if err = cursor.Decode(&data); err != nil {
		return nil, err
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
