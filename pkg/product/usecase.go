package product

import (
	"context"
	"procurement-be/utils"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProductUsecase interface {
	Store(ctx context.Context, request []ProductDataInput) (*ProductFetchResponse, error)
	Update(ctx context.Context, request []ProductDataInput) (*ProductFetchResponse, error)
	Fetch(ctx context.Context, request ProductDataInput) (*ProductFetchResponse, error)
}

type ProductUsecase struct {
	repository IProductRepository
}

func NewProductUsecase(repository IProductRepository) IProductUsecase {
	return &ProductUsecase{
		repository: repository,
	}
}

func (u *ProductUsecase) Store(ctx context.Context, request []ProductDataInput) (response *ProductFetchResponse, err error) {
	var (
		id        string
		documents []interface{}
		token     = ctx.Value("Authorization").(string)
	)
	userID, _ := utils.GetIDByToken(token)

	for i := range request {
		id = uuid.NewString()
		request[i].ID = id
		request[i].CreatedAt = time.Now().Format(utils.DatetimeLayout)
		request[i].CreatedBy = userID
		request[i].UpdatedAt = time.Now().Format(utils.DatetimeLayout)
		request[i].UpdatedBy = userID
		request[i].DeletedAt = nil

		documents = append(documents, request[i])
	}

	err = u.repository.Store(ctx, documents)
	if err != nil {
		return nil, err
	}

	response = &ProductFetchResponse{
		IsSuccess: true,
		Message:   "request successfull",
		Status:    1,
	}

	return
}

func (u *ProductUsecase) Update(ctx context.Context, request []ProductDataInput) (response *ProductFetchResponse, err error) {
	var (
		args  = []mongo.WriteModel{}
		token = ctx.Value("Authorization").(string)
	)
	userID, _ := utils.GetIDByToken(token)

	for _, v := range request {
		update := bson.M{}
		if v.Name != "" {
			update["name"] = v.Name
		}

		if v.Desc != "" {
			update["desc"] = v.Desc
		}

		if v.Price != 0 {
			update["price"] = v.Price
		}

		if v.VendorID != "" {
			update["user_id"] = v.VendorID
		}

		if v.CategoryID != "" {
			update["category_id"] = v.CategoryID
		}

		if v.ID != "" {
			update["updated_at"] = time.Now().Format(utils.DatetimeLayout)
			update["updated_by"] = userID
		}

		args = append(args, mongo.NewUpdateOneModel().
			SetFilter(bson.M{"id": v.ID}).
			SetUpdate(bson.M{"$set": update}))
	}

	err = u.repository.Update(ctx, args)
	if err != nil {
		return nil, err
	}

	response = &ProductFetchResponse{
		IsSuccess: true,
		Message:   "request successfull",
		Status:    1,
	}

	return
}

func (u *ProductUsecase) Fetch(ctx context.Context, request ProductDataInput) (response *ProductFetchResponse, err error) {
	var (
		data       []*ProductData
		filterBson = bson.M{
			"deleted_at": nil,
		}
		// token = ctx.Value("Authorization").(string)
	)
	// userID, _ := utils.GetIDByToken(token)

	if request.Name != "" {
		filterBson["name"] = primitive.Regex{Pattern: request.Name, Options: "i"}
	}

	if request.VendorID != "" {
		filterBson["user_id"] = request.VendorID
	}

	if request.CategoryID != "" {
		filterBson["category_id"] = request.CategoryID
	}

	data, err = u.repository.Fetch(ctx, filterBson)
	if err != nil {
		return nil, err
	}

	response = &ProductFetchResponse{
		IsSuccess: true,
		Message:   "request successfull",
		Status:    0,
		Data: &ProductItems{
			Items: data,
		},
	}

	return response, nil
}
