package repo

import (
	"github.com/anmimos/delivery/storage/entity"
	"github.com/golang/protobuf/ptypes/empty"
)

type CourierRepoI interface {
	CreateCourier(req *entity.CreateCourierRequest) (resp *entity.CreateCourierResponse, err error)
	GetCourier(req *entity.GetCourierRequest) (resp *entity.GetCourierResponse, err error)
	ListCouriers(req *entity.ListCourierRequest) (resp *entity.ListCourierResponse, err error)
	UpdateCourier(req *entity.UpdateCourierRequest) (resp *entity.UpdateCourierResponse, err error)
	DeleteCourier(req *entity.DeleteCourierRequest) (resp *empty.Empty, err error)
	LoginCourier(req *entity.LoginCourierRequest) (resp *entity.LoginCourierResponse, err error)
	ListCourierOrders(req *entity.ListCourierOrdersRequest) (resp *entity.ListCourierOrdersResponse, err error)
}
