package role

import (
	"context"
	"procurement-be/utils"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IRoleUsecase interface {
	Store(ctx context.Context, request []RoleData) (*RoleFetchResponse, error)
	Update(ctx context.Context, request []RoleData) (*RoleFetchResponse, error)
	Fetch(ctx context.Context, request RoleData) (*RoleFetchResponse, error)
}

type RoleUsecase struct {
	repository IRoleRepository
}

func NewRoleUsecase(repository IRoleRepository) IRoleUsecase {
	return &RoleUsecase{
		repository: repository,
	}
}

func (u *RoleUsecase) Store(ctx context.Context, request []RoleData) (response *RoleFetchResponse, err error) {
	var (
		id        string
		documents []interface{}
	)

	for i := range request {
		id = uuid.NewString()
		request[i].ID = id
		request[i].CreatedAt = time.Now().Format(utils.DatetimeLayout)
		request[i].UpdatedAt = time.Now().Format(utils.DatetimeLayout)
		request[i].DeletedAt = nil

		documents = append(documents, request[i])
	}

	err = u.repository.Store(ctx, documents)
	if err != nil {
		return nil, err
	}

	response = &RoleFetchResponse{
		IsSuccess: true,
		Message:   "request successfull",
		Status:    1,
	}

	return
}

func (u *RoleUsecase) Update(ctx context.Context, request []RoleData) (response *RoleFetchResponse, err error) {
	var (
		args = []mongo.WriteModel{}
	)

	for _, v := range request {
		update := bson.M{}
		if v.Name != "" {
			update["name"] = v.Name
			update["updated_at"] = time.Now().Format(utils.DatetimeLayout)
		}

		args = append(args, mongo.NewUpdateOneModel().
			SetFilter(bson.M{"id": v.ID}).
			SetUpdate(bson.M{"$set": update}))
	}

	err = u.repository.Update(ctx, args)
	if err != nil {
		return nil, err
	}

	response = &RoleFetchResponse{
		IsSuccess: true,
		Message:   "request successfull",
		Status:    1,
	}

	return
}

func (u *RoleUsecase) Fetch(ctx context.Context, request RoleData) (response *RoleFetchResponse, err error) {
	var (
		data       []*RoleData
		filterBson = bson.M{
			"deleted_at": nil,
		}
	)

	data, err = u.repository.Fetch(ctx, filterBson)
	if err != nil {
		return nil, err
	}

	response = &RoleFetchResponse{
		IsSuccess: true,
		Message:   "request successfull",
		Status:    0,
		Data: &RoleItems{
			Items: data,
		},
	}

	return response, nil
}
