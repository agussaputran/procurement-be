package product

type ProductData struct {
	ID        string       `json:"id,omitempty" bson:"id, omitempty"`
	Name      string       `json:"name,omitempty" bson:"name, omitempty"`
	Desc      string       `json:"desc,omitempty" bson:"desc, omitempty"`
	Vendor    VendorData   `json:"vendor,omitempty" bson:"vendor, omitempty"`
	Price     float64      `json:"price,omitempty" bson:"price, omitempty"`
	Category  CategoryData `json:"category,omitempty" bson:"category, omitempty"`
	CreatedAt string       `json:"created_at,omitempty" bson:"created_at, omitempty"`
	CreatedBy string       `json:"created_by,omitempty" bson:"created_by, omitempty"`
	UpdatedAt string       `json:"updated_at,omitempty" bson:"updated_at, omitempty"`
	UpdatedBy string       `json:"updated_by,omitempty" bson:"updated_by, omitempty"`
	DeletedAt *string      `json:"deleted_at,omitempty" bson:"deleted_at"`
}

type CategoryData struct {
	ID        string  `json:"id,omitempty" bson:"id, omitempty"`
	Name      string  `json:"name,omitempty" bson:"name, omitempty"`
	CreatedAt string  `json:"created_at,omitempty" bson:"created_at, omitempty"`
	UpdatedAt string  `json:"updated_at,omitempty" bson:"updated_at, omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty" bson:"deleted_at"`
}

type VendorData struct {
	ID        string `json:"id,omitempty" bson:"id, omitempty"`
	Role      string `json:"role,omitempty" bson:"role, omitempty"`
	Email     string `json:"email,omitempty" bson:"email, omitempty"`
	Name      string `json:"name,omitempty" bson:"name, omitempty"`
	CreatedAt string `json:"created_at,omitempty" bson:"created_at, omitempty"`
}

type ProductItems struct {
	Items []*ProductData `json:"items,omitempty"`
}

type ProductFetchResponse struct {
	IsSuccess bool          `json:"isSuccess,omitempty"`
	Message   string        `json:"message,omitempty"`
	Data      *ProductItems `json:"data,omitempty"`
	Status    int           `json:"status,omitempty"`
}

type ProductDataInput struct {
	ID         string  `json:"id,omitempty" bson:"id, omitempty"`
	Name       string  `json:"name,omitempty" bson:"name, omitempty"`
	Desc       string  `json:"desc,omitempty" bson:"desc, omitempty"`
	VendorID   string  `json:"user_id,omitempty" bson:"user_id, omitempty"`
	Price      float64 `json:"price,omitempty" bson:"price, omitempty"`
	CategoryID string  `json:"category_id,omitempty" bson:"category_id, omitempty"`
	CreatedAt  string  `json:"created_at,omitempty" bson:"created_at, omitempty"`
	CreatedBy  string  `json:"created_by,omitempty" bson:"created_by, omitempty"`
	UpdatedAt  string  `json:"updated_at,omitempty" bson:"updated_at, omitempty"`
	UpdatedBy  string  `json:"updated_by,omitempty" bson:"updated_by, omitempty"`
	DeletedAt  *string `json:"deleted_at,omitempty" bson:"deleted_at"`
}
