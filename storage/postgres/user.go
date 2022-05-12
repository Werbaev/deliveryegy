package postgres

import (
	"database/sql"

	"github.com/anmimos/delivery/storage/entity"
	"github.com/anmimos/delivery/storage/repo"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) repo.UserRepoI {
	return userRepo{db: db}
}

func (u userRepo) CreateUser(req *entity.CreateUserRequest) (resp *entity.CreateUserResponse, err error) {
	resp = &entity.CreateUserResponse{}

	query := `insert into users(name, login, password, phone_number) 
				VALUES ($1, $2, $3, $4) returning guid`

	err = u.db.Get(&resp.Guid, query, req.Name, req.Login, req.Password, req.PhoneNumber)
	if err != nil {
		return
	}

	return
}

func (u userRepo) GetUser(req *entity.GetUserRequest) (resp *entity.GetUserResponse, err error) {
	resp = &entity.GetUserResponse{}

	var result []entity.GetUserResponse

	query := `select guid, name, login, password, created_at, phone_number from users where guid = $1`

	err = u.db.Select(&result, query, req.Guid)
	if err != nil {
		return
	}

	for _, v := range result {
		resp.Guid = v.Guid
		resp.Name = v.Name
		resp.Login = v.Login
		resp.Password = v.Password
		resp.CreatedAt = v.CreatedAt
		resp.PhoneNumber = v.PhoneNumber
	}

	return
}

func (u userRepo) ListUsers(req *entity.ListUserRequest) (resp *entity.ListUserResponse, err error) {
	resp = &entity.ListUserResponse{}

	query := `select guid, name, login, password, created_at, phone_number from users order by created_at desc`

	err = u.db.Select(&resp.Users, query)
	if err != nil {
		return
	}

	return
}

func (u userRepo) UpdateUser(req *entity.UpdateUserRequest) (resp *entity.UpdateUserResponse, err error) {
	resp = &entity.UpdateUserResponse{}

	query := `update users set name = $1, login = $2, password = $3, phone_number = $4 where guid = $5`

	_, err = u.db.Exec(query, req.Name, req.Login, req.Password, req.PhoneNumber, req.Guid)
	if err != nil {
		return
	}

	resp.Guid = req.Guid

	return
}

func (u userRepo) DeleteUser(req *entity.DeleteUserRequest) (resp *empty.Empty, err error) {
	resp = &empty.Empty{}

	query := `delete from users where guid = $1`

	_, err = u.db.Exec(query, req.Guid)
	if err != nil {
		return
	}

	return
}

func (u userRepo) LoginUser(req *entity.LoginUserRequest) (resp *entity.LoginUserResponse, err error) {
	resp = &entity.LoginUserResponse{}

	var guid entity.Users

	query := `select guid, name, phone_number from users where login = $1 and password = $2`

	err = u.db.Get(&guid, query, req.Login, req.Password)

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
		resp.PhoneNumber = guid.PhoneNumber
	}

	return
}

func (u userRepo) GetUserOrders(req *entity.GetUserOrdersRequest) (resp *entity.GetUserOrdersResponse, err error) {
	resp = &entity.GetUserOrdersResponse{}

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
				where u2.guid = $1 order by o.created_at desc`

	err = u.db.Select(&result, query, req.Guid)
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

			err = u.db.Select(&resultProduct, query, pq.Array(ids))
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
			DeliveryType:  v.DeliveryType,
			PaymentType:   v.PaymentType,
			BranchName:    v.BranchName,
			MerchantName:  v.MerchantName,
		})
	}

	return
}
