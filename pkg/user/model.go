package user

type UserData struct {
	ID             string  `json:"id,omitempty" bson:"id, omitempty"`
	Role           string  `json:"role,omitempty" bson:"role, omitempty"`
	Email          string  `json:"email,omitempty" bson:"email, omitempty"`
	Name           string  `json:"name,omitempty" bson:"name, omitempty"`
	HashedPassword string  `json:"hashed_password,omitempty" bson:"hashed_password, omitempty"`
	Password       string  `json:"password,omitempty" bson:"password, omitempty"`
	CreatedAt      string  `json:"created_at,omitempty" bson:"created_at, omitempty"`
	CreatedBy      string  `json:"created_by,omitempty" bson:"created_by, omitempty"`
	UpdatedAt      string  `json:"updated_at,omitempty" bson:"updated_at, omitempty"`
	UpdatedBy      string  `json:"updated_by,omitempty" bson:"updated_by, omitempty"`
	DeletedAt      *string `json:"deleted_at,omitempty" bson:"deleted_at"`
}

type UserItems struct {
	Items []*UserData `json:"items,omitempty"`
}

type UserFetchResponse struct {
	IsSuccess bool      `json:"isSuccess,omitempty"`
	Message   string    `json:"message,omitempty"`
	Data      UserItems `json:"data,omitempty"`
	Status    int       `json:"status,omitempty"`
}

type LoginData struct {
	Token    string    `json:"token,omitempty"`
	UserData *UserData `json:"userData,omitempty"`
}

type LoginResponse struct {
	IsSuccess bool       `json:"isSuccess,omitempty"`
	Message   string     `json:"message,omitempty"`
	Data      *LoginData `json:"data,omitempty"`
	Status    int        `json:"status,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email,omitempty" bson:"email, omitempty"`
	Password string `json:"password,omitempty" bson:"password, omitempty"`
}
