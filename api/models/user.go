package models

type User struct {
	Name        string `json:"name,omitempty"`
	Login       string `json:"login,omitempty"`
	Password    string `json:"password,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type VendorUser struct {
	Name             string `json:"name,omitempty"`
	Login            string `json:"login,omitempty"`
	Password         string `json:"password,omitempty"`
	MerchantBranchId string `json:"merchant_branch_id,omitempty"`
}
