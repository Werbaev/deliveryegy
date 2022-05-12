package entity

type Merchant struct {
	CreatedAt string `json:"created_at" db:"created_at"`
	Name      string `json:"name" db:"name"`
}

type CreateMerchantRequest struct {
	Name            string  `json:"name" db:"name"`
	Logo            string  `json:"logo" db:"logo"`
	BackgroundImage string  `json:"background_image" db:"background_image"`
	Comission       float64 `json:"comission" db:"comission"`
	Status          bool    `json:"status" db:"status"`
	DeliveryTime    float64 `json:"delivery_time" db:"delivery_time"`
	Description     string  `json:"description" db:"description"`
}

type CreateMerchantResponse struct {
	Guid string `json:"guid" db:"guid"`
}

type GetMerchantRequest struct {
	Guid string `json:"guid" db:"guid"`
}

type GetMerchantResponse struct {
	Guid            string  `json:"guid" db:"guid"`
	Name            string  `json:"name" db:"name"`
	Logo            string  `json:"logo" db:"logo"`
	BackgroundImage string  `json:"background_image" db:"background_image"`
	Comission       float64 `json:"comission" db:"comission"`
	Status          bool    `json:"status" db:"status"`
	DeliveryTime    float64 `json:"delivery_time" db:"delivery_time"`
	CreatedAt       string  `json:"created_at" db:"created_at"`
	Description     string  `json:"description" db:"description"`
}

type UpdateMerchantRequest struct {
	Guid            string  `json:"guid" db:"guid"`
	Name            string  `json:"name" db:"name"`
	Logo            string  `json:"logo" db:"logo"`
	BackgroundImage string  `json:"background_image" db:"background_image"`
	Comission       float64 `json:"comission" db:"comission"`
	Status          bool    `json:"status" db:"status"`
	DeliveryTime    float64 `json:"delivery_time" db:"delivery_time"`
	Description     string  `json:"description" db:"description"`
}

type UpdateMerchantResponse struct {
	Guid string `json:"guid" db:"guid"`
}

type DeleteMerchantRequest struct {
	Guid string `json:"guid" db:"guid"`
}

type ListMerchantRequest struct {
}

type ListMerchantResponse struct {
	Merchants []GetMerchantResponse `json:"merchants"`
}

type GetMerchantBranchesRequest struct {
	Guid string `json:"guid" db:"guid"`
}

type GetMerchantBranchesResponse struct {
	MerchantsBranches []GetMerchantBranchResponse `json:"merchant_branches"`
}

type GetMerchantProductsRequest struct {
	Guid string `json:"guid" db:"guid"`
}

type GetMerchantProductsResponse struct {
	Products []GetProductResponse `json:"products"`
}

type GetMerchantOrdersRequest struct {
	MerchantId string `json:"merchant_id" db:"merchant_id"`
	Status     string `json:"status" db:"status"`
	BranchName string `json:"branch_name" db:"branch_name"`
	UserName   string `json:"user_name" db:"user_name"`
}

type GetMerchantOrdersResponse struct {
	Orders []GetOrderResponse `json:"orders"`
}
