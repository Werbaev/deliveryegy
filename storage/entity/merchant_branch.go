package entity

type MerchantBranch struct {
	CreatedAt string `json:"created_at" db:"created_at"`
	Name      string `json:"name" db:"name"`
	Address   string `json:"address" db:"address"`
	ProductId string `json:"product_id" db:"product_id"`
	UserId    string `json:"user_id" db:"user_id"`
}

type CreateMerchantBranchRequest struct {
	Name       string `json:"name" db:"name"`
	Address    string `json:"address" db:"address"`
	MerchantId string `json:"merchant_id" db:"merchant_id"`
}

type CreateMerchantBranchResponse struct {
	Guid string `json:"guid" db:"guid"`
}

type GetMerchantBranchRequest struct {
	Guid string `json:"guid" db:"guid"`
}

type GetMerchantBranchResponse struct {
	Guid       string `json:"guid" db:"guid"`
	Name       string `json:"name" db:"name"`
	Address    string `json:"address" db:"address"`
	MerchantId string `json:"merchant_id" db:"merchant_id"`
	CreatedAt  string `json:"created_at" db:"created_at"`
}

type UpdateMerchantBranchRequest struct {
	Guid       string `json:"guid" db:"guid"`
	Name       string `json:"name" db:"name"`
	Address    string `json:"address" db:"address"`
	MerchantId string `json:"merchant_id" db:"merchant_id"`
}

type UpdateMerchantBranchResponse struct {
	Guid string `json:"guid" db:"guid"`
}

type DeleteMerchantBranchRequest struct {
	Guid string `json:"guid" db:"guid"`
}

type ListMerchantBranchesRequest struct {
}

type ListMerchantBranchesResponse struct {
	MerchantsBranches []GetMerchantBranchResponse `json:"merchant_branches"`
}

type GetMerchantBranchOrdersRequest struct {
	BranchId string `json:"branch_id" db:"branch_id"`
}

type GetMerchantBranchOrdersResponse struct {
	Orders []GetOrderResponse `json:"orders"`
}
