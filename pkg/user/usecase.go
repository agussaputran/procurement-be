package user

import (
	"context"
	"os"
	"procurement-be/utils"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserUsecase interface {
	Store(ctx context.Context, request UserData) (*LoginResponse, error)
	Update(ctx context.Context, request UserData) (*LoginResponse, error)
	Fetch(ctx context.Context, request UserData) (*UserFetchResponse, error)
	Login(ctx context.Context, request LoginRequest) (*LoginResponse, error)
}

type UserUsecase struct {
	repository IUserRepository
}

func NewUserUsecase(repository IUserRepository) IUserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}

func (u *UserUsecase) Store(ctx context.Context, request UserData) (response *LoginResponse, err error) {
	var (
		user  = new(UserData)
		id    = uuid.NewString()
		token = ctx.Value("Authorization").(string)
	)
	userID, _ := utils.GetIDByToken(token)

	user, err = u.repository.FetchByEmail(ctx, request.Email)
	if err != nil && err.Error() != utils.ErrDataNotFound {
		return nil, err
	}

	if user != nil {
		response = &LoginResponse{
			IsSuccess: false,
			Message:   "Email already registered",
			Status:    1,
		}
		return
	}

	hashedPassword := utils.HashAndSalt([]byte(request.Password))
	request.HashedPassword = hashedPassword

	args := bson.M{
		"id":              id,
		"name":            request.Name,
		"email":           request.Email,
		"hashed_password": request.HashedPassword,
		"role":            request.Role,
		"timestamp":       time.Now(),
		"created_at":      time.Now().Format(utils.DatetimeLayout),
		"created_by":      userID,
		"updated_at":      time.Now().Format(utils.DatetimeLayout),
		"updated_by":      userID,
		"deleted_at":      nil,
	}
	err = u.repository.Store(ctx, args)
	if err != nil {
		return nil, err
	}

	response = &LoginResponse{
		IsSuccess: true,
		Message:   "request successfull",
		Status:    1,
	}

	return
}

func (u *UserUsecase) Update(ctx context.Context, request UserData) (response *LoginResponse, err error) {
	var (
		hashedPassword string
		token          = ctx.Value("Authorization").(string)
	)
	userID, _ := utils.GetIDByToken(token)

	if request.Password != "" {
		hashedPassword = utils.HashAndSalt([]byte(request.Password))
		request.HashedPassword = hashedPassword
	}

	args := bson.M{
		"updated_at": time.Now().Format(utils.DatetimeLayout),
	}
	if request.Name != "" {
		args["name"] = request.Name
	}
	if request.Email != "" {
		args["email"] = request.Email
	}
	if request.Role != "" {
		args["role"] = request.Role
	}
	if request.Password != "" {
		args["hashed_password"] = request.HashedPassword
	}

	filter := bson.M{
		"id": userID,
	}
	err = u.repository.Update(ctx, filter, args)
	if err != nil {
		return nil, err
	}

	response = &LoginResponse{
		IsSuccess: true,
		Message:   "request successfull",
		Status:    1,
	}

	return
}

func (u *UserUsecase) Fetch(ctx context.Context, request UserData) (response *UserFetchResponse, err error) {
	var (
		data       []*UserData
		filterBson = bson.M{
			"deleted_at": nil,
		}
	)

	if request.Name != "" {
		filterBson["name"] = primitive.Regex{Pattern: request.Name, Options: "i"}
	}

	if request.Role != "" {
		filterBson["role"] = request.Role
	}

	data, err = u.repository.Fetch(ctx, filterBson)
	if err != nil {
		return nil, err
	}

	response = &UserFetchResponse{
		IsSuccess: true,
		Message:   "request successfull",
		Status:    0,
		Data: UserItems{
			Items: data,
		},
	}

	return response, nil
}

func (u *UserUsecase) Login(ctx context.Context, request LoginRequest) (response *LoginResponse, err error) {
	var (
		data  = new(UserData)
		token string
	)
	data, err = u.repository.FetchByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	if utils.ComparePasswords(data.HashedPassword, []byte(request.Password)) {
		token, err = utils.CreateTokenJWT(utils.User{
			ID:    data.ID,
			Email: data.Email,
			Role:  data.Role,
		},
			os.Getenv("ACCESS_TOKEN_SECRET"), utils.ConvStringToInt(os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR")))
		if err != nil {
			return nil, err
		}

		response = &LoginResponse{
			IsSuccess: true,
			Message:   "Login successful",
			Status:    0,
			Data: &LoginData{
				Token: token,
				UserData: &UserData{
					ID:        data.ID,
					Role:      data.Role,
					Email:     data.Email,
					Name:      data.Name,
					CreatedAt: data.CreatedAt,
				},
			},
		}
	} else {
		response = &LoginResponse{
			IsSuccess: false,
			Message:   "Invalid username or password",
			Status:    1,
			Data:      nil,
		}
	}

	return response, nil
}
