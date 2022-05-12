package entity

type VendorUser struct {
	Guid             string `json:"guid" db:"guid"`
	CreatedAt        string `json:"created_at" db:"created_at"`
	Name             string `json:"name" db:"name"`
	Login            string `json:"login" db:"login"`
	Password         string `json:"password" db:"password"`
	MerchantBranchId string `json:"merchant_branch_id" db:"merchant_branch_id"`
}

type CreateVendorUserRequest struct {
	Name             string `json:"name" db:"name"`
	Login            string `json:"login" db:"login"`
	Password         string `json:"password" db:"password"`
	MerchantBranchId string `json:"merchant_branch_id" db:"merchant_branch_id"`
}

type CreateVendorUserResponse struct {
	Guid string `json:"guid" db:"guid"`
}

type GetVendorUserRequest struct {
	Guid string `json:"guid" db:"guid"`
}

type GetVendorUserResponse struct {
	Guid             string `json:"guid" db:"guid"`
	Name             string `json:"name" db:"name"`
	Login            string `json:"login" db:"login"`
	Password         string `json:"password" db:"password"`
	CreatedAt        string `json:"created_at" db:"created_at"`
	MerchantBranchId string `json:"merchant_branch_id" db:"merchant_branch_id"`
}

type UpdateVendorUserRequest struct {
	Name             string `json:"name" db:"name"`
	Login            string `json:"login" db:"login"`
	Password         string `json:"password" db:"password"`
	Guid             string `json:"guid" db:"guid"`
	MerchantBranchId string `json:"merchant_branch_id" db:"merchant_branch_id"`
}

type UpdateVendorUserResponse struct {
	Guid string `json:"guid" db:"guid"`
}

type DeleteVendorUserRequest struct {
	Guid string `json:"guid" db:"guid"`
}

type ListVendorUserRequest struct {
}

type ListVendorUserResponse struct {
	VendorUsers []GetVendorUserResponse `json:"vendor_users"`
}

type LoginVendorUserRequest struct {
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
}

type LoginVendorUserResponse struct {
	Exist      bool   `json:"exist"`
	Guid       string `json:"guid" db:"guid"`
	Name       string `json:"name" db:"name"`
	MerchantId string `json:"merchant_id" db:"merchant_id"`
}
