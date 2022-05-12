package entity

import (
	"database/sql"

	"github.com/jackc/pgtype"
)

type Order struct {
	CreatedAt string `json:"created_at" db:"created_at"`
	Comment   string `json:"comment" db:"comment"`
	Address   string `json:"address" db:"address"`
	ProductId string `json:"product_id" db:"product_id"`
	UserId    string `json:"user_id" db:"user_id"`
	Status    string `json:"status" db:"status"`
}

type OrderProduct struct {
	ProductId int32 `json:"product_id" db:"product_id"`
	Count     int32 `json:"count" db:"count"`
}

type CreateOrderRequest struct {
	UserId       string         `json:"user_id" db:"user_id"`
	Products     []OrderProduct `json:"products" db:"products"`
	Comment      string         `json:"comment" db:"comment"`
	Address      string         `json:"address" db:"address"`
	BranchId     string         `json:"branch_id" db:"branch_id"`
	PaymentType  string         `json:"payment_type" db:"payment_type"`
	DeliveryType string         `json:"delivery_type" db:"delivery_type"`
}

type CreateOrderResponse struct {
	Guid string `json:"guid" db:"guid"`
}

type GetOrderRequest struct {
	Guid string `json:"guid" db:"guid"`
}

type OrderProducts struct {
	Id         int32    `json:"id" db:"id"`
	Guid       string   `json:"guid" db:"guid"`
	Name       string   `json:"name" db:"name"`
	CreatedAt  string   `json:"created_at" db:"created_at"`
	Price      float64  `json:"price" db:"price"`
	CategoryId string   `json:"category_id" db:"category_id"`
	Option     []Option `json:"options" db:"options"`
	Image      string   `json:"image" db:"image"`
	Count      int32    `json:"count" db:"count"`
}

type GetOrderResponse struct {
	Guid          string          `json:"guid" db:"guid"`
	UserId        string          `json:"user_id" db:"user_id"`
	Products      []OrderProducts `json:"products" db:"products"`
	Comment       string          `json:"comment" db:"comment"`
	Address       string          `json:"address" db:"address"`
	CreatedAt     string          `json:"created_at" db:"created_at"`
	Status        string          `json:"status" db:"status"`
	BranchAddress string          `json:"branch_address" db:"branch_address"`
	BranchId      string          `json:"branch_id" db:"branch_id"`
	Comission     float64         `json:"comission" db:"comission"`
	DeliveryPrice float64         `json:"delivery_price" db:"delivery_price"`
	Image         string          `json:"image" db:"image"`
	DeliveryTime  float64         `json:"delivery_time" db:"delivery_time"`
	CourierId     string          `json:"courier_id" db:"courier_id"`
	PaymentType   string          `json:"payment_type" db:"payment_type"`
	DeliveryType  string          `json:"delivery_type" db:"delivery_type"`
	BranchName    string          `json:"branch_name" db:"branch_name"`
	MerchantName  string          `json:"merchant_name" db:"merchant_name"`
}

type GetOrderResponseStorage struct {
	Guid          string         `json:"guid" db:"guid"`
	UserId        string         `json:"user_id" db:"user_id"`
	Products      pgtype.JSONB   `json:"products" db:"products"`
	Comment       string         `json:"comment" db:"comment"`
	Address       string         `json:"address" db:"address"`
	CreatedAt     string         `json:"created_at" db:"created_at"`
	Status        string         `json:"status" db:"status"`
	BranchAddress string         `json:"branch_address" db:"branch_address"`
	BranchId      string         `json:"branch_id" db:"branch_id"`
	Comission     float64        `json:"comission" db:"comission"`
	DeliveryPrice float64        `json:"delivery_price" db:"delivery_price"`
	Image         string         `json:"image" db:"image"`
	DeliveryTime  float64        `json:"delivery_time" db:"delivery_time"`
	CourierId     sql.NullString `json:"courier_id" db:"courier_id"`
	PaymentType   string         `json:"payment_type" db:"payment_type"`
	DeliveryType  string         `json:"delivery_type" db:"delivery_type"`
	BranchName    string         `json:"branch_name" db:"branch_name"`
	MerchantName  string         `json:"merchant_name" db:"merchant_name"`
}

type UpdateOrderRequest struct {
	Guid         string         `json:"guid" db:"guid"`
	UserId       string         `json:"user_id" db:"user_id"`
	Comment      string         `json:"comment" db:"comment"`
	Address      string         `json:"address" db:"address"`
	Status       string         `json:"status" db:"status"`
	BranchId     string         `json:"branch_id" db:"branch_id"`
	PaymentType  string         `json:"payment_type" db:"payment_type"`
	DeliveryType string         `json:"delivery_type" db:"delivery_type"`
	Products     []OrderProduct `json:"products" db:"products"`
}

type UpdateOrderResponse struct {
	Guid string `json:"guid" db:"guid"`
}

type DeleteOrderRequest struct {
	Guid string `json:"guid" db:"guid"`
}

type ListOrderRequest struct {
	Status       string `json:"status" db:"status"`
	MerchantName string `json:"merchant_name" db:"merchant_name"`
	BranchName   string `json:"branch_name" db:"branch_name"`
	UserName     string `json:"user_name" db:"user_name"`
}

type ListOrderResponse struct {
	Orders []GetOrderResponse `json:"orders"`
}

type UpdateOrderStatusRequest struct {
	Guid      string `json:"guid" db:"guid"`
	Status    string `json:"status" db:"status"`
	CourierId string `json:"courier_id" db:"courier_id"`
}
