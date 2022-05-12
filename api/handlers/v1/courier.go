package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/werbaev/deliveryegy/api/models"
	"github.com/werbaev/deliveryegy/storage/entity"
)

// @Router /v1/courier/create [POST]
// @Summary Create courier
// @Description API that create courier
// @Tags courier
// @Accept  json
// @Produce  json
// @Param body body models.User true "body"
// @Success 200 {object} entity.CreateCourierResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateCourier(c *gin.Context) {
	data := models.User{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.CourierRepo().CreateCourier(&entity.CreateCourierRequest{
		Name:     data.Name,
		Login:    data.Login,
		Password: data.Password,
	})

	if h.handleError(c, err, "Error create courier") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/courier/get/{guid} [GET]
// @Summary Get courier
// @Description API that get courier
// @Tags courier
// @Accept  json
// @Produce  json
// @Param guid path string true "guid"
// @Success 200 {object} entity.GetCourierResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCourier(c *gin.Context) {
	response, err := h.storage.CourierRepo().GetCourier(&entity.GetCourierRequest{
		Guid: c.Param("guid"),
	})

	if h.handleError(c, err, "Error get courier") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/courier/list [GET]
// @Summary List couriers
// @Description API list couriers
// @Tags courier
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.ListCourierResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ListCouriers(c *gin.Context) {
	response, err := h.storage.CourierRepo().ListCouriers(&entity.ListCourierRequest{})

	if h.handleError(c, err, "Error get list couriers") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/courier/{guid}/list-orders [GET]
// @Summary List courier orders
// @Description API list courier orders
// @Tags courier
// @Accept  json
// @Produce  json
// @Param guid path string true "guid"
// @Success 200 {object} entity.ListCourierOrdersResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ListCourierOrders(c *gin.Context) {
	response, err := h.storage.CourierRepo().ListCourierOrders(&entity.ListCourierOrdersRequest{
		CourierId: c.Param("guid"),
	})

	if h.handleError(c, err, "Error get list courier orders") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/courier/update [PUT]
// @Summary Update courier
// @Description API that update courier
// @Tags courier
// @Accept  json
// @Produce  json
// @Param body body entity.UpdateCourierRequest true "body"
// @Success 200 {object} entity.UpdateCourierResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateCourier(c *gin.Context) {
	data := entity.UpdateCourierRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.CourierRepo().UpdateCourier(&entity.UpdateCourierRequest{
		Name:     data.Name,
		Login:    data.Login,
		Password: data.Password,
		Guid:     data.Guid,
	})

	if h.handleError(c, err, "Error update courier") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/courier/delete [DELETE]
// @Summary Delete courier
// @Description API that delete courier
// @Tags courier
// @Accept  json
// @Produce  json
// @Param body body entity.DeleteCourierRequest true "body"
// @Success 200 {object} entity.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteCourier(c *gin.Context) {
	data := entity.DeleteCourierRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.CourierRepo().DeleteCourier(&entity.DeleteCourierRequest{
		Guid: data.Guid,
	})

	if h.handleError(c, err, "Error delete courier") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/courier/login [POST]
// @Summary Login courier
// @Description Login courier
// @Tags courier
// @Accept  json
// @Produce  json
// @Param body body entity.LoginCourierRequest true "body"
// @Success 200 {object} entity.LoginCourierResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) LoginCourier(c *gin.Context) {
	data := entity.LoginCourierRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.CourierRepo().LoginCourier(&entity.LoginCourierRequest{
		Login:    data.Login,
		Password: data.Password,
	})

	if h.handleError(c, err, "Error login courier") {
		return
	}

	c.JSON(200, response)
}
