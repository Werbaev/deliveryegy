package postgres

import (
	"database/sql"

	"github.com/anmimos/delivery/storage/entity"
	"github.com/anmimos/delivery/storage/repo"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type courierRepo struct {
	db *sqlx.DB
}

func NewCourierRepo(db *sqlx.DB) repo.CourierRepoI {
	return courierRepo{db: db}
}

func (c courierRepo) CreateCourier(req *entity.CreateCourierRequest) (resp *entity.CreateCourierResponse, err error) {
	resp = &entity.CreateCourierResponse{}

	query := `insert into couriers(name, login, password) 
				VALUES ($1, $2, $3) returning guid`

	err = c.db.Get(&resp.Guid, query, req.Name, req.Login, req.Password)
	if err != nil {
		return
	}

	return
}

func (c courierRepo) GetCourier(req *entity.GetCourierRequest) (resp *entity.GetCourierResponse, err error) {
	resp = &entity.GetCourierResponse{}

	var result []entity.GetCourierResponse

	query := `select guid, name, login, password, created_at from couriers where guid = $1`

	err = c.db.Select(&result, query, req.Guid)
	if err != nil {
		return
	}

	for _, v := range result {
		resp.Guid = v.Guid
		resp.Name = v.Name
		resp.Login = v.Login
		resp.Password = v.Password
		resp.CreatedAt = v.CreatedAt
	}

	return
}

func (c courierRepo) ListCouriers(req *entity.ListCourierRequest) (resp *entity.ListCourierResponse, err error) {
	resp = &entity.ListCourierResponse{}

	query := `select guid, name, login, password, created_at from couriers order by created_at desc`

	err = c.db.Select(&resp.Couriers, query)
	if err != nil {
		return
	}

	return
}

func (c courierRepo) ListCourierOrders(req *entity.ListCourierOrdersRequest) (resp *entity.ListCourierOrdersResponse, err error) {
	resp = &entity.ListCourierOrdersResponse{}

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
				where o.courier_id = (select id from couriers c where c.guid = $1) order by o.created_at desc`

	err = c.db.Select(&result, query, req.CourierId)
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

			err = c.db.Select(&resultProduct, query, pq.Array(ids))
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

func (c courierRepo) UpdateCourier(req *entity.UpdateCourierRequest) (resp *entity.UpdateCourierResponse, err error) {
	resp = &entity.UpdateCourierResponse{}

	query := `update couriers set name = $1, login = $2, password = $3 where guid = $4`

	_, err = c.db.Exec(query, req.Name, req.Login, req.Password, req.Guid)
	if err != nil {
		return
	}

	resp.Guid = req.Guid

	return
}

func (c courierRepo) DeleteCourier(req *entity.DeleteCourierRequest) (resp *empty.Empty, err error) {
	resp = &empty.Empty{}

	query := `delete from couriers where guid = $1`

	_, err = c.db.Exec(query, req.Guid)
	if err != nil {
		return
	}

	return
}

func (c courierRepo) LoginCourier(req *entity.LoginCourierRequest) (resp *entity.LoginCourierResponse, err error) {
	resp = &entity.LoginCourierResponse{}

	var guid entity.Users

	query := `select guid, name from couriers where login = $1 and password = $2`

	err = c.db.Get(&guid, query, req.Login, req.Password)

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
	}

	return
}
