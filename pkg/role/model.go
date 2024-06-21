package role

type RoleData struct {
	ID        string  `json:"id,omitempty" bson:"id, omitempty"`
	Name      string  `json:"name,omitempty" bson:"name, omitempty"`
	CreatedAt string  `json:"created_at,omitempty" bson:"created_at, omitempty"`
	CreatedBy string  `json:"created_by,omitempty" bson:"created_by, omitempty"`
	UpdatedAt string  `json:"updated_at,omitempty" bson:"updated_at, omitempty"`
	UpdatedBy string  `json:"updated_by,omitempty" bson:"updated_by, omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty" bson:"deleted_at"`
}

type RoleItems struct {
	Items []*RoleData `json:"items,omitempty"`
}

type RoleFetchResponse struct {
	IsSuccess bool       `json:"isSuccess,omitempty"`
	Message   string     `json:"message,omitempty"`
	Data      *RoleItems `json:"data,omitempty"`
	Status    int        `json:"status,omitempty"`
}
