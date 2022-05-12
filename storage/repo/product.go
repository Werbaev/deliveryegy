package repo

import (
	"github.com/anmimos/delivery/storage/entity"
	"github.com/golang/protobuf/ptypes/empty"
)

type ProductRepoI interface {
	CreateProduct(req *entity.CreateProductRequest) (resp *entity.CreateProductResponse, err error)
	GetProduct(req *entity.GetProductRequest) (resp *entity.GetProductResponse, err error)
	ListProducts(req *entity.ListProductRequest) (resp *entity.ListProductResponse, err error)
	UpdateProduct(req *entity.UpdateProductRequest) (resp *entity.UpdateProductResponse, err error)
	DeleteProduct(req *entity.DeleteProductRequest) (resp *empty.Empty, err error)
}
