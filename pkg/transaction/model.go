package transaction

type TransactionData struct {
	ID          string       `json:"id,omitempty" bson:"id, omitempty"`
	Product     ProductData  `json:"product,omitempty" bson:"product, omitempty"`
	Vendor      VendorData   `json:"vendor,omitempty" bson:"vendor, omitempty"`
	TotalPrice  float64      `json:"total_price,omitempty" bson:"total_price, omitempty"`
	Category    CategoryData `json:"category,omitempty" bson:"category, omitempty"`
	CreatedAt   string       `json:"created_at,omitempty" bson:"created_at, omitempty"`
	CreatedBy   string       `json:"created_by,omitempty" bson:"created_by, omitempty"`
	UpdatedAt   string       `json:"updated_at,omitempty" bson:"updated_at, omitempty"`
	UpdatedBy   string       `json:"updated_by,omitempty" bson:"updated_by, omitempty"`
	RequestedBy string       `json:"requested_by,omitempty" bson:"requested_by, omitempty"`
	ApprovedBy  string       `json:"approved_by,omitempty" bson:"approved_by, omitempty"`
	DeletedAt   *string      `json:"deleted_at,omitempty" bson:"deleted_at"`
}

type StatusProduct struct {
	Code      string `json:"code,omitempty" bson:"code, omitempty"`
	Message   string `json:"message,omitempty" bson:"message, omitempty"`
	CreatedAt string `json:"created_at,omitempty" bson:"created_at, omitempty"`
}

type ProductData struct {
	ID     string          `json:"id,omitempty" bson:"id, omitempty"`
	Name   string          `json:"name,omitempty" bson:"name, omitempty"`
	Desc   string          `json:"desc,omitempty" bson:"desc, omitempty"`
	Vendor VendorData      `json:"vendor,omitempty" bson:"vendor, omitempty"`
	Qty    int64           `json:"qty,omitempty" bson:"qty, omitempty"`
	Price  float64         `json:"price,omitempty" bson:"price, omitempty"`
	Status []StatusProduct `json:"status,omitempty" bson:"status, omitempty"`
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

type TransactionItems struct {
	Items []*TransactionData `json:"items,omitempty"`
}

type TransactionFetchResponse struct {
	IsSuccess bool              `json:"isSuccess,omitempty"`
	Message   string            `json:"message,omitempty"`
	Data      *TransactionItems `json:"data,omitempty"`
	Status    int               `json:"status,omitempty"`
}

type TransactionDataInput struct {
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
