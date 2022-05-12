package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/werbaev/deliveryegy/storage/postgres"
	"github.com/werbaev/deliveryegy/storage/repo"
)

type StorageI interface {
	UserRepo() repo.UserRepoI
	CategoryRepo() repo.CategoryRepoI
	ProductRepo() repo.ProductRepoI
	OrderRepo() repo.OrderRepoI
	MerchantRepo() repo.MerchantRepoI
	MerchantBranchRepo() repo.MerchantBranchRepoI
	CourierRepo() repo.CourierRepoI
	VendorUserRepo() repo.VendorUserRepoI
}

type storagePg struct {
	db                 *sqlx.DB
	userRepo           repo.UserRepoI
	categoryRepo       repo.CategoryRepoI
	productRepo        repo.ProductRepoI
	orderRepo          repo.OrderRepoI
	merchantRepo       repo.MerchantRepoI
	merchantBranchRepo repo.MerchantBranchRepoI
	courierRepo        repo.CourierRepoI
	vendorUserRepo     repo.VendorUserRepoI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db:                 db,
		userRepo:           postgres.NewUserRepo(db),
		categoryRepo:       postgres.NewCategoryRepo(db),
		productRepo:        postgres.NewProductRepo(db),
		orderRepo:          postgres.NewOrderRepo(db),
		merchantRepo:       postgres.NewMerchantRepo(db),
		merchantBranchRepo: postgres.NewMerchantBranchRepo(db),
		courierRepo:        postgres.NewCourierRepo(db),
		vendorUserRepo:     postgres.NewVendorUserRepo(db),
	}
}

func (s storagePg) UserRepo() repo.UserRepoI {
	return s.userRepo
}

func (s storagePg) CategoryRepo() repo.CategoryRepoI {
	return s.categoryRepo
}

func (s storagePg) ProductRepo() repo.ProductRepoI {
	return s.productRepo
}

func (s storagePg) OrderRepo() repo.OrderRepoI {
	return s.orderRepo
}

func (s storagePg) MerchantRepo() repo.MerchantRepoI {
	return s.merchantRepo
}

func (s storagePg) MerchantBranchRepo() repo.MerchantBranchRepoI {
	return s.merchantBranchRepo
}

func (s storagePg) CourierRepo() repo.CourierRepoI {
	return s.courierRepo
}

func (s storagePg) VendorUserRepo() repo.VendorUserRepoI {
	return s.vendorUserRepo
}
