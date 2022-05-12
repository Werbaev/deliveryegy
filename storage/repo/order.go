package repo

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/werbaev/deliveryegy/storage/entity"
)

type OrderRepoI interface {
	CreateOrder(req *entity.CreateOrderRequest) (resp *entity.CreateOrderResponse, err error)
	GetOrder(req *entity.GetOrderRequest) (resp *entity.GetOrderResponse, err error)
	ListOrders(req *entity.ListOrderRequest) (resp *entity.ListOrderResponse, err error)
	UpdateOrder(req *entity.UpdateOrderRequest) (resp *entity.UpdateOrderResponse, err error)
	UpdateOrderStatus(req *entity.UpdateOrderStatusRequest) (resp *entity.UpdateOrderResponse, err error)
	DeleteOrder(req *entity.DeleteOrderRequest) (resp *empty.Empty, err error)
}
