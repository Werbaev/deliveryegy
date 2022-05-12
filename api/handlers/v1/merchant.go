package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/werbaev/deliveryegy/storage/entity"
)

// @Router /v1/merchant/create [POST]
// @Summary Create merchant
// @Description API that create merchant
// @Tags merchant
// @Accept  json
// @Produce  json
// @Param body body entity.CreateMerchantRequest true "body"
// @Success 200 {object} entity.CreateMerchantResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateMerchant(c *gin.Context) {
	data := entity.CreateMerchantRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.MerchantRepo().CreateMerchant(&entity.CreateMerchantRequest{
		Name:            data.Name,
		Logo:            data.Logo,
		BackgroundImage: data.BackgroundImage,
		Comission:       data.Comission,
		Status:          data.Status,
		DeliveryTime:    data.DeliveryTime,
		Description:     data.Description,
	})

	if h.handleError(c, err, "Error create merchant") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/merchant/get/{guid} [GET]
// @Summary Get merchant
// @Description API that get merchant
// @Tags merchant
// @Accept  json
// @Produce  json
// @Param guid path string true "guid"
// @Success 200 {object} entity.GetMerchantResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetMerchant(c *gin.Context) {
	response, err := h.storage.MerchantRepo().GetMerchant(&entity.GetMerchantRequest{
		Guid: c.Param("guid"),
	})

	if h.handleError(c, err, "Error get merchant") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/merchant/list [GET]
// @Summary List merchants
// @Description API list merchants
// @Tags merchant
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.ListMerchantResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ListMerchants(c *gin.Context) {
	response, err := h.storage.MerchantRepo().ListMerchants(&entity.ListMerchantRequest{})

	if h.handleError(c, err, "Error get list merchants") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/merchant/update [PUT]
// @Summary Update merchant
// @Description API that update merchant
// @Tags merchant
// @Accept  json
// @Produce  json
// @Param body body entity.UpdateMerchantRequest true "body"
// @Success 200 {object} entity.UpdateMerchantResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateMerchant(c *gin.Context) {
	data := entity.UpdateMerchantRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.MerchantRepo().UpdateMerchant(&entity.UpdateMerchantRequest{
		Guid:            data.Guid,
		Name:            data.Name,
		Logo:            data.Logo,
		BackgroundImage: data.BackgroundImage,
		Comission:       data.Comission,
		Status:          data.Status,
		DeliveryTime:    data.DeliveryTime,
		Description:     data.Description,
	})

	if h.handleError(c, err, "Error update merchant") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/merchant/delete [DELETE]
// @Summary Delete merchant
// @Description API that delete merchant
// @Tags merchant
// @Accept  json
// @Produce  json
// @Param body body entity.DeleteMerchantRequest true "body"
// @Success 200 {object} entity.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteMerchant(c *gin.Context) {
	data := entity.DeleteMerchantRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.MerchantRepo().DeleteMerchant(&entity.DeleteMerchantRequest{
		Guid: data.Guid,
	})

	if h.handleError(c, err, "Error delete merchant") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/merchant/get-branches/{guid} [GET]
// @Summary Get merchant-branches by merchant id
// @Description API that get merchant-branches
// @Tags merchant
// @Accept  json
// @Produce  json
// @Param guid path string true "guid"
// @Success 200 {object} entity.GetMerchantBranchesResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetMerchantBranches(c *gin.Context) {
	response, err := h.storage.MerchantRepo().GetMerchantBranches(&entity.GetMerchantBranchesRequest{
		Guid: c.Param("guid"),
	})

	if h.handleError(c, err, "Error get merchant branches") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/merchant/products/{guid} [GET]
// @Summary Get merchant-products by merchant id
// @Description API that get merchant-products
// @Tags merchant
// @Accept  json
// @Produce  json
// @Param guid path string true "guid"
// @Success 200 {object} entity.GetMerchantProductsResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetMerchantProducts(c *gin.Context) {
	response, err := h.storage.MerchantRepo().GetMerchantProducts(&entity.GetMerchantProductsRequest{
		Guid: c.Param("guid"),
	})

	if h.handleError(c, err, "Error get merchant products") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/merchant/orders/{guid} [GET]
// @Summary Get merchant-orders by merchant id
// @Description API that get merchant-orders
// @Tags merchant
// @Accept  json
// @Produce  json
// @Param guid path string true "guid"
// @Param status query string false "status"
// @Param branch_name query string false "branch_name"
// @Param user_name query string false "user_name"
// @Success 200 {object} entity.GetMerchantOrdersResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetMerchantOrders(c *gin.Context) {
	response, err := h.storage.MerchantRepo().GetMerchantOrders(&entity.GetMerchantOrdersRequest{
		MerchantId: c.Param("guid"),
		Status:     c.Query("status"),
		BranchName: c.Query("branch_name"),
		UserName:   c.Query("user_name"),
	})

	if h.handleError(c, err, "Error get merchant orders") {
		return
	}

	c.JSON(200, response)
}
