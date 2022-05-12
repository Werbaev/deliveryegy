package postgres

import (
	"github.com/anmimos/delivery/storage/entity"
	"github.com/anmimos/delivery/storage/repo"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
)

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) repo.CategoryRepoI {
	return categoryRepo{db: db}
}

func (c categoryRepo) CreateCategory(req *entity.CreateCategoryRequest) (resp *entity.CreateCategoryResponse, err error) {
	resp = &entity.CreateCategoryResponse{}

	query := `insert into categories(name, merchant_id) 
				VALUES ($1, (select id from merchants where guid = $2)) returning guid`

	err = c.db.Get(&resp.Guid, query, req.Name, req.MerchantId)
	if err != nil {
		return
	}

	return
}

func (c categoryRepo) GetCategory(req *entity.GetCategoryRequest) (resp *entity.GetCategoryResponse, err error) {
	resp = &entity.GetCategoryResponse{}

	var result []entity.GetCategoryResponse

	query := `select guid, name, created_at, (select guid from merchants where id = merchant_id) as merchant_id from categories where guid = $1`

	err = c.db.Select(&result, query, req.Guid)
	if err != nil {
		return
	}

	for _, v := range result {
		resp.Guid = v.Guid
		resp.Name = v.Name
		resp.CreatedAt = v.CreatedAt
		resp.MerchantId = v.MerchantId
	}

	return
}

func (c categoryRepo) ListCategories(req *entity.ListCategoryRequest) (resp *entity.ListCategoryResponse, err error) {
	resp = &entity.ListCategoryResponse{}

	query := `select guid, name, created_at, (select guid from merchants where id = merchant_id) as merchant_id from categories order by created_at desc`

	err = c.db.Select(&resp.Categories, query)
	if err != nil {
		return
	}

	return
}

func (c categoryRepo) UpdateCategory(req *entity.UpdateCategoryRequest) (resp *entity.UpdateCategoryResponse, err error) {
	resp = &entity.UpdateCategoryResponse{}

	query := `update categories set name = $1, merchant_id = (select id from merchants where guid = $2) where guid = $3`

	_, err = c.db.Exec(query, req.Name, req.MerchantId, req.Guid)
	if err != nil {
		return
	}

	resp.Guid = req.Guid

	return
}

func (c categoryRepo) DeleteCategory(req *entity.DeleteCategoryRequest) (resp *empty.Empty, err error) {
	resp = &empty.Empty{}

	query := `delete from categories where guid = $1`

	_, err = c.db.Exec(query, req.Guid)
	if err != nil {
		return
	}

	return
}
