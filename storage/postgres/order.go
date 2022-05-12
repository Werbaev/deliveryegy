package postgres

import (
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/werbaev/deliveryegy/storage/entity"
	"github.com/werbaev/deliveryegy/storage/repo"
)

type orderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) repo.OrderRepoI {
	return orderRepo{db: db}
}

func (o orderRepo) CreateOrder(req *entity.CreateOrderRequest) (resp *entity.CreateOrderResponse, err error) {
	resp = &entity.CreateOrderResponse{}
	query := `insert into orders(user_id, products, comment, address, branch_id, payment_type, delivery_type) 
				VALUES ((select id from users where guid = $1), $2, $3, $4, (select id from merchant_branches where guid = $5), $6, $7) returning guid`

	err = o.db.Get(&resp.Guid, query, req.UserId, req.Products, req.Comment, req.Address, req.BranchId, req.PaymentType, req.DeliveryType)
	if err != nil {
		return
	}

	return
}

func (o orderRepo) GetOrder(req *entity.GetOrderRequest) (resp *entity.GetOrderResponse, err error) {
	resp = &entity.GetOrderResponse{}

	var (
		result        []entity.GetOrderResponseStorage
		resultProduct []entity.GetProductResponseStorage
		orderProducts []entity.OrderProduct
		ids           []int
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
				join merchants m on mb.merchant_id = m.id where o.guid = $1`

	err = o.db.Select(&result, query, req.Guid)
	if err != nil {
		return
	}

	query = `select id, guid, name, created_at, price, (select guid from categories c where c.id = category_id) as category_id, options, image from products where id = ANY($1)`

	result[0].Products.AssignTo(&orderProducts)

	for _, v := range orderProducts {
		ids = append(ids, int(v.ProductId))
	}

	err = o.db.Select(&resultProduct, query, pq.Array(ids))
	if err != nil {
		return
	}

	for _, v := range result {
		var products []entity.OrderProducts

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

		resp.Guid = v.Guid
		resp.UserId = v.UserId
		resp.Products = products
		resp.Comment = v.Comment
		resp.Address = v.Address
		resp.CreatedAt = v.CreatedAt
		resp.Status = v.Status
		resp.BranchId = v.BranchId
		resp.BranchAddress = v.BranchAddress
		resp.Comission = v.Comission
		resp.DeliveryPrice = v.DeliveryPrice
		resp.Image = v.Image
		resp.CourierId = v.CourierId.String
		resp.DeliveryTime = v.DeliveryTime
		resp.PaymentType = v.PaymentType
		resp.DeliveryType = v.DeliveryType
		resp.BranchName = v.BranchName
		resp.MerchantName = v.MerchantName
	}

	return
}

func (o orderRepo) ListOrders(req *entity.ListOrderRequest) (resp *entity.ListOrderResponse, err error) {
	resp = &entity.ListOrderResponse{
		Orders: []entity.GetOrderResponse{},
	}

	filter := ` where 1=1`

	if req.Status != "" {
		filter += fmt.Sprintf(` and o.status = '%s'`, req.Status)
	}

	if req.MerchantName != "" {
		filter += fmt.Sprintf(` and m.name ilike '%s'`, "%"+req.MerchantName+"%")
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
				join merchants m on mb.merchant_id = m.id` + filter + ` order by o.created_at desc`

	err = o.db.Select(&result, query)
	if err != nil {
		return
	}

	for i, v := range result {

		var (
			orderProducts []entity.OrderProduct
			ids           []int
			products      []entity.OrderProducts
			resultProduct []entity.GetProductResponseStorage
		)

		result[i].Products.AssignTo(&orderProducts)

		for _, v := range orderProducts {
			ids = append(ids, int(v.ProductId))
		}

		if len(orderProducts) > 0 {
			query = `select id, guid, name, created_at, price, (select guid from categories c where c.id = category_id) as category_id, options, image from products where id = ANY($1)`

			err = o.db.Select(&resultProduct, query, pq.Array(ids))
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

func (o orderRepo) UpdateOrder(req *entity.UpdateOrderRequest) (resp *entity.UpdateOrderResponse, err error) {
	resp = &entity.UpdateOrderResponse{}

	query := `update orders set user_id = (select id from users u where u.guid = $1), products = $2, comment = $3, address = $4, status = $5, branch_id = (select id from merchant_branches where guid = $6), payment_type = $7, delivery_type = $8 where guid = $9`

	_, err = o.db.Exec(query, req.UserId, req.Products, req.Comment, req.Address, req.Status, req.BranchId, req.PaymentType, req.DeliveryType, req.Guid)
	if err != nil {
		return
	}

	resp.Guid = req.Guid

	return
}

func (o orderRepo) UpdateOrderStatus(req *entity.UpdateOrderStatusRequest) (resp *entity.UpdateOrderResponse, err error) {
	resp = &entity.UpdateOrderResponse{}

	query := `update orders set status = $1, courier_id = (select id from couriers where guid = $2) where guid = $3`

	_, err = o.db.Exec(query, req.Status, req.CourierId, req.Guid)
	if err != nil {
		return
	}

	resp.Guid = req.Guid

	return
}

func (o orderRepo) DeleteOrder(req *entity.DeleteOrderRequest) (resp *empty.Empty, err error) {
	resp = &empty.Empty{}

	query := `delete from orders where guid = $1`

	_, err = o.db.Exec(query, req.Guid)
	if err != nil {
		return
	}

	return
}
