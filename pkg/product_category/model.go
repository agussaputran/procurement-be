package product_category

type ProductCategoryData struct {
	ID        string  `json:"id,omitempty" bson:"id, omitempty"`
	Name      string  `json:"name,omitempty" bson:"name, omitempty"`
	CreatedAt string  `json:"created_at,omitempty" bson:"created_at, omitempty"`
	CreatedBy string  `json:"created_by,omitempty" bson:"created_by, omitempty"`
	UpdatedAt string  `json:"updated_at,omitempty" bson:"updated_at, omitempty"`
	UpdatedBy string  `json:"updated_by,omitempty" bson:"updated_by, omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty" bson:"deleted_at"`
}

type ProductCategoryItems struct {
	Items []*ProductCategoryData `json:"items,omitempty"`
}

type ProductCategoryFetchResponse struct {
	IsSuccess bool                  `json:"isSuccess,omitempty"`
	Message   string                `json:"message,omitempty"`
	Data      *ProductCategoryItems `json:"data,omitempty"`
	Status    int                   `json:"status,omitempty"`
}
