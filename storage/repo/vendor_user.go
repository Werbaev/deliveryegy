package repo

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/werbaev/deliveryegy/storage/entity"
)

type VendorUserRepoI interface {
	CreateVendorUser(req *entity.CreateVendorUserRequest) (resp *entity.CreateVendorUserResponse, err error)
	GetVendorUser(req *entity.GetVendorUserRequest) (resp *entity.GetVendorUserResponse, err error)
	ListVendorUsers(req *entity.ListVendorUserRequest) (resp *entity.ListVendorUserResponse, err error)
	UpdateVendorUser(req *entity.UpdateVendorUserRequest) (resp *entity.UpdateVendorUserResponse, err error)
	DeleteVendorUser(req *entity.DeleteVendorUserRequest) (resp *empty.Empty, err error)
	LoginVendorUser(req *entity.LoginVendorUserRequest) (resp *entity.LoginVendorUserResponse, err error)
}
