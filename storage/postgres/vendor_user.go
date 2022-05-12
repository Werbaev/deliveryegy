package postgres

import (
	"database/sql"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"github.com/werbaev/deliveryegy/storage/entity"
	"github.com/werbaev/deliveryegy/storage/repo"
)

type vendorUserRepo struct {
	db *sqlx.DB
}

func NewVendorUserRepo(db *sqlx.DB) repo.VendorUserRepoI {
	return vendorUserRepo{db: db}
}

func (v vendorUserRepo) CreateVendorUser(req *entity.CreateVendorUserRequest) (resp *entity.CreateVendorUserResponse, err error) {
	resp = &entity.CreateVendorUserResponse{}

	query := `insert into vendor_users(name, login, password, merchant_branch_id) 
				VALUES ($1, $2, $3, (select id from merchant_branches where guid = $4)) returning guid`

	err = v.db.Get(&resp.Guid, query, req.Name, req.Login, req.Password, req.MerchantBranchId)
	if err != nil {
		return
	}

	return
}

func (v vendorUserRepo) GetVendorUser(req *entity.GetVendorUserRequest) (resp *entity.GetVendorUserResponse, err error) {
	resp = &entity.GetVendorUserResponse{}

	var result []entity.GetVendorUserResponse

	query := `select guid, name, login, password, created_at, (select guid from merchant_branches where id = merchant_branch_id) as merchant_branch_id from vendor_users where guid = $1`

	err = v.db.Select(&result, query, req.Guid)
	if err != nil {
		return
	}

	for _, v := range result {
		resp.Guid = v.Guid
		resp.Name = v.Name
		resp.Login = v.Login
		resp.Password = v.Password
		resp.CreatedAt = v.CreatedAt
		resp.MerchantBranchId = v.MerchantBranchId
	}

	return
}

func (v vendorUserRepo) ListVendorUsers(req *entity.ListVendorUserRequest) (resp *entity.ListVendorUserResponse, err error) {
	resp = &entity.ListVendorUserResponse{}

	query := `select guid, name, login, password, created_at, (select guid from merchant_branches where id = merchant_branch_id) as merchant_branch_id from vendor_users order by created_at desc`

	err = v.db.Select(&resp.VendorUsers, query)
	if err != nil {
		return
	}

	return
}

func (v vendorUserRepo) UpdateVendorUser(req *entity.UpdateVendorUserRequest) (resp *entity.UpdateVendorUserResponse, err error) {
	resp = &entity.UpdateVendorUserResponse{}

	query := `update vendor_users set name = $1, login = $2, password = $3, merchant_branch_id = (select id from merchant_branches where guid = $4) where guid = $5`

	_, err = v.db.Exec(query, req.Name, req.Login, req.Password, req.MerchantBranchId, req.Guid)
	if err != nil {
		return
	}

	resp.Guid = req.Guid

	return
}

func (v vendorUserRepo) DeleteVendorUser(req *entity.DeleteVendorUserRequest) (resp *empty.Empty, err error) {
	resp = &empty.Empty{}

	query := `delete from vendor_users where guid = $1`

	_, err = v.db.Exec(query, req.Guid)
	if err != nil {
		return
	}

	return
}

func (v vendorUserRepo) LoginVendorUser(req *entity.LoginVendorUserRequest) (resp *entity.LoginVendorUserResponse, err error) {
	resp = &entity.LoginVendorUserResponse{}

	var guid entity.LoginVendorUserResponse

	query := `select 
				v.guid, 
				v.name, 
				m.guid as merchant_id 
			from vendor_users v 
			join merchant_branches mb on mb.id = v.merchant_branch_id 
			join merchants m on m.id = mb.merchant_id 
			where login = $1 and password = $2`

	err = v.db.Get(&guid, query, req.Login, req.Password)

	if err == sql.ErrNoRows {
		return resp, nil
	}

	if err != nil {
		return
	}

	if guid.Guid != "" {
		resp.Exist = true
		resp.Guid = guid.Guid
		resp.Name = guid.Name
		resp.MerchantId = guid.MerchantId
	}

	return
}
