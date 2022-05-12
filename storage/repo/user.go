package repo

import (
	"github.com/anmimos/delivery/storage/entity"
	"github.com/golang/protobuf/ptypes/empty"
)

type UserRepoI interface {
	CreateUser(req *entity.CreateUserRequest) (resp *entity.CreateUserResponse, err error)
	GetUser(req *entity.GetUserRequest) (resp *entity.GetUserResponse, err error)
	ListUsers(req *entity.ListUserRequest) (resp *entity.ListUserResponse, err error)
	UpdateUser(req *entity.UpdateUserRequest) (resp *entity.UpdateUserResponse, err error)
	DeleteUser(req *entity.DeleteUserRequest) (resp *empty.Empty, err error)
	LoginUser(req *entity.LoginUserRequest) (resp *entity.LoginUserResponse, err error)
	GetUserOrders(req *entity.GetUserOrdersRequest) (resp *entity.GetUserOrdersResponse, err error)
}
