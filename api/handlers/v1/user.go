package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/werbaev/deliveryegy/api/models"
	"github.com/werbaev/deliveryegy/storage/entity"
)

// @Router /v1/user/create [POST]
// @Summary Create user
// @Description API that create user
// @Tags user
// @Accept  json
// @Produce  json
// @Param body body models.User true "body"
// @Success 200 {object} entity.CreateUserResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateUser(c *gin.Context) {
	data := models.User{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.UserRepo().CreateUser(&entity.CreateUserRequest{
		Name:        data.Name,
		Login:       data.Login,
		Password:    data.Password,
		PhoneNumber: data.PhoneNumber,
	})

	if h.handleError(c, err, "Error create user") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/user/get/{guid} [GET]
// @Summary Get user
// @Description API that get user
// @Tags user
// @Accept  json
// @Produce  json
// @Param guid path string true "guid"
// @Success 200 {object} entity.GetUserResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetUser(c *gin.Context) {
	response, err := h.storage.UserRepo().GetUser(&entity.GetUserRequest{
		Guid: c.Param("guid"),
	})

	if h.handleError(c, err, "Error get user") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/user/list [GET]
// @Summary List users
// @Description API list users
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.ListUserResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ListUsers(c *gin.Context) {
	response, err := h.storage.UserRepo().ListUsers(&entity.ListUserRequest{})

	if h.handleError(c, err, "Error get list users") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/user/update [PUT]
// @Summary Update user
// @Description API that update user
// @Tags user
// @Accept  json
// @Produce  json
// @Param body body entity.UpdateUserRequest true "body"
// @Success 200 {object} entity.UpdateUserResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateUser(c *gin.Context) {
	data := entity.UpdateUserRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.UserRepo().UpdateUser(&entity.UpdateUserRequest{
		Name:        data.Name,
		Login:       data.Login,
		Password:    data.Password,
		Guid:        data.Guid,
		PhoneNumber: data.PhoneNumber,
	})

	if h.handleError(c, err, "Error update user") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/user/delete [DELETE]
// @Summary Delete user
// @Description API that delete user
// @Tags user
// @Accept  json
// @Produce  json
// @Param body body entity.DeleteUserRequest true "body"
// @Success 200 {object} entity.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteUser(c *gin.Context) {
	data := entity.DeleteUserRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.UserRepo().DeleteUser(&entity.DeleteUserRequest{
		Guid: data.Guid,
	})

	if h.handleError(c, err, "Error delete user") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/user/login [POST]
// @Summary Login user
// @Description Login user
// @Tags user
// @Accept  json
// @Produce  json
// @Param body body entity.LoginUserRequest true "body"
// @Success 200 {object} entity.LoginUserResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) LoginUser(c *gin.Context) {
	data := entity.LoginUserRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.UserRepo().LoginUser(&entity.LoginUserRequest{
		Login:    data.Login,
		Password: data.Password,
	})

	if h.handleError(c, err, "Error login user") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/user/{guid}/orders [GET]
// @Summary Get user orders
// @Description API that get user orders
// @Tags user
// @Accept  json
// @Produce  json
// @Param guid path string true "guid"
// @Success 200 {object} entity.GetUserOrdersResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetUserOrders(c *gin.Context) {
	response, err := h.storage.UserRepo().GetUserOrders(&entity.GetUserOrdersRequest{
		Guid: c.Param("guid"),
	})

	if h.handleError(c, err, "Error get user orders") {
		return
	}

	c.JSON(200, response)
}
