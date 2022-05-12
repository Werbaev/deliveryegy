package repo

import (
	"github.com/anmimos/delivery/storage/entity"
	"github.com/golang/protobuf/ptypes/empty"
)

type MerchantBranchRepoI interface {
	CreateMerchantBranch(req *entity.CreateMerchantBranchRequest) (resp *entity.CreateMerchantBranchResponse, err error)
	GetMerchantBranch(req *entity.GetMerchantBranchRequest) (resp *entity.GetMerchantBranchResponse, err error)
	ListMerchantBranches(req *entity.ListMerchantBranchesRequest) (resp *entity.ListMerchantBranchesResponse, err error)
	UpdateMerchantBranch(req *entity.UpdateMerchantBranchRequest) (resp *entity.UpdateMerchantBranchResponse, err error)
	DeleteMerchantBranch(req *entity.DeleteMerchantBranchRequest) (resp *empty.Empty, err error)
	GetMerchantBranchOrders(req *entity.GetMerchantBranchOrdersRequest) (resp *entity.GetMerchantBranchOrdersResponse, err error)
}
