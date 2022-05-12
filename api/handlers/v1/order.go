package v1

import (
	"github.com/anmimos/delivery/storage/entity"
	"github.com/gin-gonic/gin"
)

// @Router /v1/order/create [POST]
// @Summary Create order
// @Description API that create order
// @Tags order
// @Accept  json
// @Produce  json
// @Param body body entity.CreateOrderRequest true "body"
// @Success 200 {object} entity.CreateOrderResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateOrder(c *gin.Context) {
	data := entity.CreateOrderRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.OrderRepo().CreateOrder(&entity.CreateOrderRequest{
		UserId:       data.UserId,
		Products:     data.Products,
		Comment:      data.Comment,
		Address:      data.Address,
		BranchId:     data.BranchId,
		PaymentType:  data.PaymentType,
		DeliveryType: data.DeliveryType,
	})

	if h.handleError(c, err, "Error create order") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/order/get/{guid} [GET]
// @Summary Get order
// @Description API that get order
// @Tags order
// @Accept  json
// @Produce  json
// @Param guid path string true "guid"
// @Success 200 {object} entity.GetOrderResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetOrder(c *gin.Context) {
	response, err := h.storage.OrderRepo().GetOrder(&entity.GetOrderRequest{
		Guid: c.Param("guid"),
	})

	if h.handleError(c, err, "Error get order") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/order/list [GET]
// @Summary List orders
// @Description API list orders
// @Tags order
// @Accept  json
// @Produce  json
// @Param status query string false "status"
// @Param merchant_name query string false "merchant_name"
// @Param branch_name query string false "branch_name"
// @Param user_name query string false "user_name"
// @Success 200 {object} entity.ListOrderResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ListOrders(c *gin.Context) {
	response, err := h.storage.OrderRepo().ListOrders(&entity.ListOrderRequest{
		Status:       c.Query("status"),
		MerchantName: c.Query("merchant_name"),
		BranchName:   c.Query("branch_name"),
		UserName:     c.Query("user_name"),
	})

	if h.handleError(c, err, "Error get list orders") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/order/update [PUT]
// @Summary Update order
// @Description API that update order
// @Tags order
// @Accept  json
// @Produce  json
// @Param body body entity.UpdateOrderRequest true "body"
// @Success 200 {object} entity.UpdateOrderResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateOrder(c *gin.Context) {
	data := entity.UpdateOrderRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.OrderRepo().UpdateOrder(&entity.UpdateOrderRequest{
		Guid:         data.Guid,
		UserId:       data.UserId,
		Products:     data.Products,
		Comment:      data.Comment,
		Address:      data.Address,
		Status:       data.Status,
		BranchId:     data.BranchId,
		PaymentType:  data.PaymentType,
		DeliveryType: data.DeliveryType,
	})

	if h.handleError(c, err, "Error update order") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/order/update-status [PUT]
// @Summary Update order status
// @Description API that update order status
// @Tags order
// @Accept  json
// @Produce  json
// @Param body body entity.UpdateOrderStatusRequest true "body"
// @Success 200 {object} entity.UpdateOrderResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateOrderStatus(c *gin.Context) {
	data := entity.UpdateOrderStatusRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.OrderRepo().UpdateOrderStatus(&entity.UpdateOrderStatusRequest{
		Guid:      data.Guid,
		Status:    data.Status,
		CourierId: data.CourierId,
	})

	if h.handleError(c, err, "Error update order status") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/order/delete [DELETE]
// @Summary Delete order
// @Description API that delete order
// @Tags order
// @Accept  json
// @Produce  json
// @Param body body entity.DeleteOrderRequest true "body"
// @Success 200 {object} entity.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteOrder(c *gin.Context) {
	data := entity.DeleteOrderRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.OrderRepo().DeleteOrder(&entity.DeleteOrderRequest{
		Guid: data.Guid,
	})

	if h.handleError(c, err, "Error delete order") {
		return
	}

	c.JSON(200, response)
}
