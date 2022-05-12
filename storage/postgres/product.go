package postgres

import (
	"github.com/anmimos/delivery/storage/entity"
	"github.com/anmimos/delivery/storage/repo"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
)

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) repo.ProductRepoI {
	return productRepo{db: db}
}

func (p productRepo) CreateProduct(req *entity.CreateProductRequest) (resp *entity.CreateProductResponse, err error) {
	resp = &entity.CreateProductResponse{}

	query := `insert into products(name, price, category_id, options, image) 
				VALUES ($1, $2, (select id from categories where guid = $3), $4, $5) returning guid`

	err = p.db.Get(&resp.Guid, query, req.Name, req.Price, req.CategoryId, req.Option, req.Image)
	if err != nil {
		return
	}

	return
}

func (p productRepo) GetProduct(req *entity.GetProductRequest) (resp *entity.GetProductResponse, err error) {
	resp = &entity.GetProductResponse{}

	var (
		result []entity.GetProductResponseStorage
	)

	query := `select id, guid, name, created_at, price, (select guid from categories c where c.id = category_id) as category_id, options, image from products where guid = $1`

	err = p.db.Select(&result, query, req.Guid)
	if err != nil {
		return
	}

	for _, v := range result {
		resp.Id = v.Id
		resp.Guid = v.Guid
		resp.Name = v.Name
		resp.CreatedAt = v.CreatedAt
		resp.Price = v.Price
		resp.CategoryId = v.CategoryId
		_ = v.Option.AssignTo(&resp.Option)
		resp.Image = v.Image.String
	}

	return
}

func (p productRepo) ListProducts(req *entity.ListProductRequest) (resp *entity.ListProductResponse, err error) {
	resp = &entity.ListProductResponse{}

	var result []entity.GetProductResponseStorage

	query := `select id, guid, name, created_at, price, (select guid from categories where id = category_id) as category_id, options, image from products order by created_at desc`

	err = p.db.Select(&result, query)
	if err != nil {
		return
	}

	for _, v := range result {
		var option entity.GetProductResponse
		_ = v.Option.AssignTo(&option.Option)

		resp.Products = append(resp.Products, entity.GetProductResponse{
			Id:         v.Id,
			Guid:       v.Guid,
			Name:       v.Name,
			CreatedAt:  v.CreatedAt,
			Price:      v.Price,
			CategoryId: v.CategoryId,
			Option:     option.Option,
			Image:      v.Image.String,
		})
	}

	return
}

func (p productRepo) UpdateProduct(req *entity.UpdateProductRequest) (resp *entity.UpdateProductResponse, err error) {
	resp = &entity.UpdateProductResponse{}

	query := `update products set name = $1, price = $2, category_id = (select id from categories c where c.guid = $3), options = $4, image = $5 where guid = $6`

	_, err = p.db.Exec(query, req.Name, req.Price, req.CategoryId, req.Option, req.Image, req.Guid)
	if err != nil {
		return
	}

	resp.Guid = req.Guid

	return
}

func (p productRepo) DeleteProduct(req *entity.DeleteProductRequest) (resp *empty.Empty, err error) {
	resp = &empty.Empty{}

	query := `delete from products where guid = $1`

	_, err = p.db.Exec(query, req.Guid)
	if err != nil {
		return
	}

	return
}
