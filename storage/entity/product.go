package entity

import (
	"database/sql"

	"github.com/jackc/pgtype"
)

type Product struct {
	CreatedAt  string  `json:"created_at" db:"created_at"`
	Name       string  `json:"name" db:"name"`
	Price      float64 `json:"price" db:"price"`
	CategoryId string  `json:"category_id" db:"category_id"`
}

type Option struct {
	Name  string  `json:"name" db:"name"`
	Price float64 `json:"price" db:"price"`
}

type CreateProductRequest struct {
	Name       string   `json:"name" db:"name"`
	Price      float64  `json:"price" db:"price"`
	CategoryId string   `json:"category_id" db:"category_id"`
	Option     []Option `json:"options" db:"options"`
	Image      string   `json:"image" db:"image"`
}

type CreateProductResponse struct {
	Guid string `json:"guid" db:"guid"`
}

type GetProductRequest struct {
	Guid string `json:"guid" db:"guid"`
}

type GetProductResponseStorage struct {
	Id         int32          `json:"id" db:"id"`
	Guid       string         `json:"guid" db:"guid"`
	Name       string         `json:"name" db:"name"`
	CreatedAt  string         `json:"created_at" db:"created_at"`
	Price      float64        `json:"price" db:"price"`
	CategoryId string         `json:"category_id" db:"category_id"`
	Option     pgtype.JSONB   `json:"options" db:"options"`
	Image      sql.NullString `json:"image" db:"image"`
}

type GetProductResponse struct {
	Id         int32    `json:"id" db:"id"`
	Guid       string   `json:"guid" db:"guid"`
	Name       string   `json:"name" db:"name"`
	CreatedAt  string   `json:"created_at" db:"created_at"`
	Price      float64  `json:"price" db:"price"`
	CategoryId string   `json:"category_id" db:"category_id"`
	Option     []Option `json:"options" db:"options"`
	Image      string   `json:"image" db:"image"`
}

type UpdateProductRequest struct {
	Name       string   `json:"name" db:"name"`
	Guid       string   `json:"guid" db:"guid"`
	Price      float64  `json:"price" db:"price"`
	CategoryId string   `json:"category_id" db:"category_id"`
	Option     []Option `json:"options" db:"options"`
	Image      string   `json:"image" db:"image"`
}

type UpdateProductResponse struct {
	Guid string `json:"guid" db:"guid"`
}

type DeleteProductRequest struct {
	Guid string `json:"guid" db:"guid"`
}

type ListProductRequest struct {
}

type ListProductResponse struct {
	Products []GetProductResponse `json:"products"`
}
