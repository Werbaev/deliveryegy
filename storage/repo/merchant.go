package repo

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/werbaev/deliveryegy/storage/entity"
)

type MerchantRepoI interface {
	CreateMerchant(req *entity.CreateMerchantRequest) (resp *entity.CreateMerchantResponse, err error)
	GetMerchant(req *entity.GetMerchantRequest) (resp *entity.GetMerchantResponse, err error)
	GetMerchantBranches(req *entity.GetMerchantBranchesRequest) (resp *entity.GetMerchantBranchesResponse, err error)
	ListMerchants(req *entity.ListMerchantRequest) (resp *entity.ListMerchantResponse, err error)
	UpdateMerchant(req *entity.UpdateMerchantRequest) (resp *entity.UpdateMerchantResponse, err error)
	DeleteMerchant(req *entity.DeleteMerchantRequest) (resp *empty.Empty, err error)
	GetMerchantProducts(req *entity.GetMerchantProductsRequest) (resp *entity.GetMerchantProductsResponse, err error)
	GetMerchantOrders(req *entity.GetMerchantOrdersRequest) (resp *entity.GetMerchantOrdersResponse, err error)
}
