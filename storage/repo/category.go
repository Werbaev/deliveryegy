package repo

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/werbaev/deliveryegy/storage/entity"
)

type CategoryRepoI interface {
	CreateCategory(req *entity.CreateCategoryRequest) (resp *entity.CreateCategoryResponse, err error)
	GetCategory(req *entity.GetCategoryRequest) (resp *entity.GetCategoryResponse, err error)
	ListCategories(req *entity.ListCategoryRequest) (resp *entity.ListCategoryResponse, err error)
	UpdateCategory(req *entity.UpdateCategoryRequest) (resp *entity.UpdateCategoryResponse, err error)
	DeleteCategory(req *entity.DeleteCategoryRequest) (resp *empty.Empty, err error)
}
