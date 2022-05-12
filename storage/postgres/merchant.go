package postgres

import (
	"fmt"

	"github.com/anmimos/delivery/storage/entity"
	"github.com/anmimos/delivery/storage/repo"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type merchantRepo struct {
	db *sqlx.DB
}

func NewMerchantRepo(db *sqlx.DB) repo.MerchantRepoI {
	return merchantRepo{db: db}
}

func (m merchantRepo) CreateMerchant(req *entity.CreateMerchantRequest) (resp *entity.CreateMerchantResponse, err error) {
	resp = &entity.CreateMerchantResponse{}

	query := `insert into merchants(name, logo, background_image, comission, status, delivery_time, description)
				VALUES ($1, $2, $3, $4, $5, $6, $7) returning guid`

	err = m.db.Get(&resp.Guid, query, req.Name, req.Logo, req.BackgroundImage, req.Comission, req.Status, req.DeliveryTime, req.Description)
	if err != nil {
		return
	}

	return
}

func (m merchantRepo) GetMerchant(req *entity.GetMerchantRequest) (resp *entity.GetMerchantResponse, err error) {
	resp = &entity.GetMerchantResponse{}

	var result []entity.GetMerchantResponse

	query := `select guid, name, logo, background_image, comission, status, delivery_time, created_at, description from merchants where guid = $1`

	err = m.db.Select(&result, query, req.Guid)
	if err != nil {
		return
	}

	for _, v := range result {
		resp.Guid = v.Guid
		resp.Name = v.Name
		resp.Logo = v.Logo
		resp.BackgroundImage = v.BackgroundImage
		resp.Comission = v.Comission
		resp.Status = v.Status
		resp.DeliveryTime = v.DeliveryTime
		resp.CreatedAt = v.CreatedAt
		resp.Description = v.Description
	}

	return
}

func (m merchantRepo) ListMerchants(req *entity.ListMerchantRequest) (resp *entity.ListMerchantResponse, err error) {
	resp = &entity.ListMerchantResponse{}

	query := `select guid, name, logo, background_image, comission, status, delivery_time, created_at, description from merchants order by created_at desc`

	err = m.db.Select(&resp.Merchants, query)
	if err != nil {
		return
	}

	return
}

func (m merchantRepo) UpdateMerchant(req *entity.UpdateMerchantRequest) (resp *entity.UpdateMerchantResponse, err error) {
	resp = &entity.UpdateMerchantResponse{}

	query := `update merchants set name = $1, logo = $2, background_image = $3, comission = $4, status = $5, delivery_time = $6, description = $7 where guid = $8`

	_, err = m.db.Exec(query, req.Name, req.Logo, req.BackgroundImage, req.Comission, req.Status, req.DeliveryTime, req.Description, req.Guid)
	if err != nil {
		return
	}

	resp.Guid = req.Guid

	return
}

func (m merchantRepo) DeleteMerchant(req *entity.DeleteMerchantRequest) (resp *empty.Empty, err error) {
	resp = &empty.Empty{}

	query := `delete from merchants where guid = $1`

	_, err = m.db.Exec(query, req.Guid)
	if err != nil {
		return
	}

	return
}

func (m merchantRepo) GetMerchantBranches(req *entity.GetMerchantBranchesRequest) (resp *entity.GetMerchantBranchesResponse, err error) {
	resp = &entity.GetMerchantBranchesResponse{}

	query := `select guid, name, address, (select guid from merchants where id = merchant_id) as merchant_id, created_at from merchant_branches where merchant_id = (select id from merchants where guid = $1) order by created_at desc`

	err = m.db.Select(&resp.MerchantsBranches, query, req.Guid)
	if err != nil {
		return
	}

	return
}

func (m merchantRepo) GetMerchantProducts(req *entity.GetMerchantProductsRequest) (resp *entity.GetMerchantProductsResponse, err error) {
	resp = &entity.GetMerchantProductsResponse{}

	var result []entity.GetProductResponseStorage

	query := `select
				id, 
				guid, 
				name, 
				created_at, 
				price, 
				(select guid from categories where id = category_id) as category_id, 
				options, 
				image
 			from products where category_id in (select id from categories where merchant_id = (select id from merchants m where m.guid = $1)) order by created_at desc`

	err = m.db.Select(&result, query, req.Guid)
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

func (m merchantRepo) GetMerchantOrders(req *entity.GetMerchantOrdersRequest) (resp *entity.GetMerchantOrdersResponse, err error) {
	resp = &entity.GetMerchantOrdersResponse{}

	filter := ` and 1=1`

	if req.Status != "" {
		filter += fmt.Sprintf(` and o.status = '%s'`, req.Status)
	}

	if req.BranchName != "" {
		filter += fmt.Sprintf(` and mb.name ilike '%s'`, "%"+req.BranchName+"%")
	}

	if req.UserName != "" {
		filter += fmt.Sprintf(` and u2.name ilike '%s'`, "%"+req.UserName+"%")
	}

	var (
		result []entity.GetOrderResponseStorage
	)

	query := `select o.guid,
					u2.guid as user_id,
					comment,
					o.address,
					o.created_at,
					o.status,
					mb.guid as branch_id,
					mb.address as branch_address,
					m.comission,
					m.delivery_time,
					o.products,
					(select guid from couriers c where c.id = o.courier_id) as courier_id,
					o.delivery_price,
					o.payment_type,
					o.delivery_type,
					mb.name as branch_name,
					m.name as merchant_name
				from orders o
				join users u2 on u2.id = o.user_id
				join merchant_branches mb on o.branch_id = mb.id
				join merchants m on mb.merchant_id = m.id
				where m.guid = $1` + filter + ` order by o.created_at desc`

	err = m.db.Select(&result, query, req.MerchantId)
	if err != nil {
		return
	}

	for i, v := range result {

		var (
			products      []entity.OrderProducts
			resultProduct []entity.GetProductResponseStorage
			orderProducts []entity.OrderProduct
			ids           []int
		)

		result[i].Products.AssignTo(&orderProducts)

		for _, v := range orderProducts {
			ids = append(ids, int(v.ProductId))
		}

		if len(orderProducts) > 0 {
			query = `select id, guid, name, created_at, price, (select guid from categories c where c.id = category_id) as category_id, options, image from products where id = ANY($1)`

			err = m.db.Select(&resultProduct, query, pq.Array(ids))
			if err != nil {
				return
			}

			for _, v2 := range resultProduct {
				var (
					option []entity.Option
					count  int32
				)

				_ = v2.Option.AssignTo(&option)

				for _, k := range orderProducts {
					if k.ProductId == v2.Id {
						count = k.Count
					}
				}

				products = append(products, entity.OrderProducts{
					Guid:       v2.Guid,
					Name:       v2.Name,
					CreatedAt:  v2.CreatedAt,
					Price:      v2.Price,
					CategoryId: v2.CategoryId,
					Image:      v2.Image.String,
					Option:     option,
					Id:         v2.Id,
					Count:      count,
				})
			}
		}

		resp.Orders = append(resp.Orders, entity.GetOrderResponse{
			Guid:          v.Guid,
			UserId:        v.UserId,
			Products:      products,
			Comment:       v.Comment,
			Address:       v.Address,
			CreatedAt:     v.CreatedAt,
			Status:        v.Status,
			BranchId:      v.BranchId,
			BranchAddress: v.BranchAddress,
			Comission:     v.Comission,
			DeliveryPrice: v.DeliveryPrice,
			Image:         v.Image,
			CourierId:     v.CourierId.String,
			DeliveryTime:  v.DeliveryTime,
			PaymentType:   v.PaymentType,
			DeliveryType:  v.DeliveryType,
			BranchName:    v.BranchName,
			MerchantName:  v.MerchantName,
		})
	}

	return
}
