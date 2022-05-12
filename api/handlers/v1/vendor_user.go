package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/werbaev/deliveryegy/api/models"
	"github.com/werbaev/deliveryegy/storage/entity"
)

// @Router /v1/vendor-user/create [POST]
// @Summary Create vendor-user
// @Description API that create vendor-user
// @Tags vendor-user
// @Accept  json
// @Produce  json
// @Param body body models.VendorUser true "body"
// @Success 200 {object} entity.CreateVendorUserResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateVendorUser(c *gin.Context) {
	data := models.VendorUser{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.VendorUserRepo().CreateVendorUser(&entity.CreateVendorUserRequest{
		Name:             data.Name,
		Login:            data.Login,
		Password:         data.Password,
		MerchantBranchId: data.MerchantBranchId,
	})

	if h.handleError(c, err, "Error create vendor-user") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/vendor-user/get/{guid} [GET]
// @Summary Get vendor-user
// @Description API that get vendor-user
// @Tags vendor-user
// @Accept  json
// @Produce  json
// @Param guid path string true "guid"
// @Success 200 {object} entity.GetVendorUserResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetVendorUser(c *gin.Context) {
	response, err := h.storage.VendorUserRepo().GetVendorUser(&entity.GetVendorUserRequest{
		Guid: c.Param("guid"),
	})

	if h.handleError(c, err, "Error get vendor-user") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/vendor-user/list [GET]
// @Summary List vendor-users
// @Description API list vendor-users
// @Tags vendor-user
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.ListVendorUserResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ListVendorUsers(c *gin.Context) {
	response, err := h.storage.VendorUserRepo().ListVendorUsers(&entity.ListVendorUserRequest{})

	if h.handleError(c, err, "Error get list vendor-users") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/vendor-user/update [PUT]
// @Summary Update vendor-user
// @Description API that update vendor-user
// @Tags vendor-user
// @Accept  json
// @Produce  json
// @Param body body entity.UpdateVendorUserRequest true "body"
// @Success 200 {object} entity.UpdateVendorUserResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateVendorUser(c *gin.Context) {
	data := entity.UpdateVendorUserRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.VendorUserRepo().UpdateVendorUser(&entity.UpdateVendorUserRequest{
		Name:             data.Name,
		Login:            data.Login,
		Password:         data.Password,
		Guid:             data.Guid,
		MerchantBranchId: data.MerchantBranchId,
	})

	if h.handleError(c, err, "Error update vendor-user") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/vendor-user/delete [DELETE]
// @Summary Delete vendor-user
// @Description API that delete vendor-user
// @Tags vendor-user
// @Accept  json
// @Produce  json
// @Param body body entity.DeleteVendorUserRequest true "body"
// @Success 200 {object} entity.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteVendorUser(c *gin.Context) {
	data := entity.DeleteVendorUserRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.VendorUserRepo().DeleteVendorUser(&entity.DeleteVendorUserRequest{
		Guid: data.Guid,
	})

	if h.handleError(c, err, "Error delete vendor-user") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/vendor-user/login [POST]
// @Summary Login vendor user
// @Description Login vendor user
// @Tags vendor-user
// @Accept  json
// @Produce  json
// @Param body body entity.LoginVendorUserRequest true "body"
// @Success 200 {object} entity.LoginVendorUserResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) LoginVendorUser(c *gin.Context) {
	data := entity.LoginVendorUserRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.VendorUserRepo().LoginVendorUser(&entity.LoginVendorUserRequest{
		Login:    data.Login,
		Password: data.Password,
	})

	if h.handleError(c, err, "Error login vendor-user") {
		return
	}

	c.JSON(200, response)
}
