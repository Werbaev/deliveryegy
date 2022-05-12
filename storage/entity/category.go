package entity

type Category struct {
	CreatedAt string `db:"created_at"`
	Name      string `db:"name"`
}

type CreateCategoryRequest struct {
	Name       string `json:"name" db:"name"`
	MerchantId string `json:"merchant_id" db:"merchant_id"`
}

type CreateCategoryResponse struct {
	Guid string `json:"guid" db:"guid"`
}

type GetCategoryRequest struct {
	Guid string `json:"guid" db:"guid"`
}

type GetCategoryResponse struct {
	Guid       string `json:"guid" db:"guid"`
	Name       string `json:"name" db:"name"`
	CreatedAt  string `json:"created_at" db:"created_at"`
	MerchantId string `json:"merchant_id" db:"merchant_id"`
}

type UpdateCategoryRequest struct {
	Name       string `json:"name" db:"name"`
	Guid       string `json:"guid" db:"guid"`
	MerchantId string `json:"merchant_id" db:"merchant_id"`
}

type UpdateCategoryResponse struct {
	Guid string `json:"guid" db:"guid"`
}

type DeleteCategoryRequest struct {
	Guid string `json:"guid" db:"guid"`
}

type ListCategoryRequest struct {
}

type ListCategoryResponse struct {
	Categories []GetCategoryResponse `json:"categories"`
}
