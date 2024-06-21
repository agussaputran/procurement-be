package product_category

import (
	"context"
	"procurement-be/utils"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProductCategoryUsecase interface {
	Store(ctx context.Context, request []ProductCategoryData) (*ProductCategoryFetchResponse, error)
	Update(ctx context.Context, request []ProductCategoryData) (*ProductCategoryFetchResponse, error)
	Fetch(ctx context.Context, request ProductCategoryData) (*ProductCategoryFetchResponse, error)
}

type ProductCategoryUsecase struct {
	repository IProductCategoryRepository
}

func NewProductCategoryUsecase(repository IProductCategoryRepository) IProductCategoryUsecase {
	return &ProductCategoryUsecase{
		repository: repository,
	}
}

func (u *ProductCategoryUsecase) Store(ctx context.Context, request []ProductCategoryData) (response *ProductCategoryFetchResponse, err error) {
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
		request[i].UpdatedAt = time.Now().Format(utils.DatetimeLayout)
		request[i].CreatedBy = userID
		request[i].UpdatedBy = userID
		request[i].DeletedAt = nil

		documents = append(documents, request[i])
	}

	err = u.repository.Store(ctx, documents)
	if err != nil {
		return nil, err
	}

	response = &ProductCategoryFetchResponse{
		IsSuccess: true,
		Message:   "request successfull",
		Status:    1,
	}

	return
}

func (u *ProductCategoryUsecase) Update(ctx context.Context, request []ProductCategoryData) (response *ProductCategoryFetchResponse, err error) {
	var (
		args  = []mongo.WriteModel{}
		token = ctx.Value("Authorization").(string)
	)
	userID, _ := utils.GetIDByToken(token)

	for _, v := range request {
		update := bson.M{}
		if v.Name != "" {
			update["name"] = v.Name
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

	response = &ProductCategoryFetchResponse{
		IsSuccess: true,
		Message:   "request successfull",
		Status:    1,
	}

	return
}

func (u *ProductCategoryUsecase) Fetch(ctx context.Context, request ProductCategoryData) (response *ProductCategoryFetchResponse, err error) {
	var (
		data       []*ProductCategoryData
		filterBson = bson.M{
			"deleted_at": nil,
		}
	)

	data, err = u.repository.Fetch(ctx, filterBson)
	if err != nil {
		return nil, err
	}

	response = &ProductCategoryFetchResponse{
		IsSuccess: true,
		Message:   "request successfull",
		Status:    0,
		Data: &ProductCategoryItems{
			Items: data,
		},
	}

	return response, nil
}
