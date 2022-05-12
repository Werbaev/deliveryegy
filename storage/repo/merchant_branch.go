package repo

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/werbaev/deliveryegy/storage/entity"
)

type MerchantBranchRepoI interface {
	CreateMerchantBranch(req *entity.CreateMerchantBranchRequest) (resp *entity.CreateMerchantBranchResponse, err error)
	GetMerchantBranch(req *entity.GetMerchantBranchRequest) (resp *entity.GetMerchantBranchResponse, err error)
	ListMerchantBranches(req *entity.ListMerchantBranchesRequest) (resp *entity.ListMerchantBranchesResponse, err error)
	UpdateMerchantBranch(req *entity.UpdateMerchantBranchRequest) (resp *entity.UpdateMerchantBranchResponse, err error)
	DeleteMerchantBranch(req *entity.DeleteMerchantBranchRequest) (resp *empty.Empty, err error)
	GetMerchantBranchOrders(req *entity.GetMerchantBranchOrdersRequest) (resp *entity.GetMerchantBranchOrdersResponse, err error)
}
