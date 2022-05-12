package postgres

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"github.com/werbaev/deliveryegy/storage/entity"
	"github.com/werbaev/deliveryegy/storage/repo"
)

type merchantBranchRepo struct {
	db *sqlx.DB
}

func NewMerchantBranchRepo(db *sqlx.DB) repo.MerchantBranchRepoI {
	return merchantBranchRepo{db: db}
}

func (m merchantBranchRepo) CreateMerchantBranch(req *entity.CreateMerchantBranchRequest) (resp *entity.CreateMerchantBranchResponse, err error) {
	resp = &entity.CreateMerchantBranchResponse{}

	query := `insert into merchant_branches(name, address, merchant_id) 
				VALUES ($1, $2, (select id from merchants where guid = $3)) returning guid`

	err = m.db.Get(&resp.Guid, query, req.Name, req.Address, req.MerchantId)
	if err != nil {
		return
	}

	return
}

func (m merchantBranchRepo) GetMerchantBranch(req *entity.GetMerchantBranchRequest) (resp *entity.GetMerchantBranchResponse, err error) {
	resp = &entity.GetMerchantBranchResponse{}

	var result []entity.GetMerchantBranchResponse

	query := `select guid, name, address, (select guid from merchants where id = merchant_id) as merchant_id, created_at from merchant_branches where guid = $1`

	err = m.db.Select(&result, query, req.Guid)
	if err != nil {
		return
	}

	for _, v := range result {
		resp.Guid = v.Guid
		resp.Name = v.Name
		resp.Address = v.Address
		resp.MerchantId = v.MerchantId
		resp.CreatedAt = v.CreatedAt
	}

	return
}

func (m merchantBranchRepo) ListMerchantBranches(req *entity.ListMerchantBranchesRequest) (resp *entity.ListMerchantBranchesResponse, err error) {
	resp = &entity.ListMerchantBranchesResponse{}

	query := `select guid, name, address, (select guid from merchants where id = merchant_id) as merchant_id, created_at from merchant_branches order by created_at desc`

	err = m.db.Select(&resp.MerchantsBranches, query)
	if err != nil {
		return
	}

	return
}

func (m merchantBranchRepo) UpdateMerchantBranch(req *entity.UpdateMerchantBranchRequest) (resp *entity.UpdateMerchantBranchResponse, err error) {
	resp = &entity.UpdateMerchantBranchResponse{}

	query := `update merchant_branches set name = $1, address = $2, merchant_id = (select id from merchants where guid = $3) where guid = $4`

	_, err = m.db.Exec(query, req.Name, req.Address, req.MerchantId, req.Guid)
	if err != nil {
		return
	}

	resp.Guid = req.Guid

	return
}

func (m merchantBranchRepo) DeleteMerchantBranch(req *entity.DeleteMerchantBranchRequest) (resp *empty.Empty, err error) {
	resp = &empty.Empty{}

	query := `delete from merchant_branches where guid = $1`

	_, err = m.db.Exec(query, req.Guid)
	if err != nil {
		return
	}

	return
}

func (m merchantBranchRepo) GetMerchantBranchOrders(req *entity.GetMerchantBranchOrdersRequest) (resp *entity.GetMerchantBranchOrdersResponse, err error) {
	resp = &entity.GetMerchantBranchOrdersResponse{}

	query := `select o.guid,
				u2.guid as user_id,
				p2.guid as product_id,
				comment,
				o.address,
				o.created_at,
				o.status,
				mb.guid as branch_id,
				mb.address as branch_address,
				m.comission,
				p2.price,
				p2.image,
				m.delivery_time
			from orders o
			join users u2 on u2.id = o.user_id
			join products p2 on p2.id = o.product_id
			join merchant_branches mb on o.branch_id = mb.id
			join merchants m on mb.merchant_id = m.id
			where mb.guid = $1
			order by o.created_at desc`

	err = m.db.Select(&resp.Orders, query, req.BranchId)
	if err != nil {
		return
	}

	return
}
