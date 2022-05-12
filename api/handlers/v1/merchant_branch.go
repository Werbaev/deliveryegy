package v1

import (
	"github.com/anmimos/delivery/storage/entity"
	"github.com/gin-gonic/gin"
)

// @Router /v1/merchant-branch/create [POST]
// @Summary Create merchant-branch
// @Description API that create merchant-branch
// @Tags merchant-branch
// @Accept  json
// @Produce  json
// @Param body body entity.CreateMerchantBranchRequest true "body"
// @Success 200 {object} entity.CreateMerchantBranchResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateMerchantBranch(c *gin.Context) {
	data := entity.CreateMerchantBranchRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.MerchantBranchRepo().CreateMerchantBranch(&entity.CreateMerchantBranchRequest{
		Name:       data.Name,
		Address:    data.Address,
		MerchantId: data.MerchantId,
	})

	if h.handleError(c, err, "Error create merchant-branch") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/merchant-branch/get/{guid} [GET]
// @Summary Get merchant-branch
// @Description API that get merchant-branch
// @Tags merchant-branch
// @Accept  json
// @Produce  json
// @Param guid path string true "guid"
// @Success 200 {object} entity.GetMerchantBranchResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetMerchantBranch(c *gin.Context) {
	response, err := h.storage.MerchantBranchRepo().GetMerchantBranch(&entity.GetMerchantBranchRequest{
		Guid: c.Param("guid"),
	})

	if h.handleError(c, err, "Error get merchant-branch") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/merchant-branch/list [GET]
// @Summary List merchant-branches
// @Description API list merchants-branches
// @Tags merchant-branch
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.ListMerchantBranchesResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ListMerchantBranches(c *gin.Context) {
	response, err := h.storage.MerchantBranchRepo().ListMerchantBranches(&entity.ListMerchantBranchesRequest{})

	if h.handleError(c, err, "Error get list merchant-branch") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/merchant-branch/update [PUT]
// @Summary Update merchant-branch
// @Description API that update merchant-branch
// @Tags merchant-branch
// @Accept  json
// @Produce  json
// @Param body body entity.UpdateMerchantBranchRequest true "body"
// @Success 200 {object} entity.UpdateMerchantBranchResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateMerchantBranch(c *gin.Context) {
	data := entity.UpdateMerchantBranchRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.MerchantBranchRepo().UpdateMerchantBranch(&entity.UpdateMerchantBranchRequest{
		Guid:       data.Guid,
		Name:       data.Name,
		Address:    data.Address,
		MerchantId: data.MerchantId,
	})

	if h.handleError(c, err, "Error update merchant-branch") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/merchant-branch/delete [DELETE]
// @Summary Delete merchant-branch
// @Description API that delete merchant-branch
// @Tags merchant-branch
// @Accept  json
// @Produce  json
// @Param body body entity.DeleteMerchantBranchRequest true "body"
// @Success 200 {object} entity.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteMerchantBranch(c *gin.Context) {
	data := entity.DeleteMerchantBranchRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.MerchantBranchRepo().DeleteMerchantBranch(&entity.DeleteMerchantBranchRequest{
		Guid: data.Guid,
	})

	if h.handleError(c, err, "Error delete merchant-branch") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/merchant-branch/{guid}/orders [GET]
// @Summary Get merchant-branch orders
// @Description API that get merchant-branch orders
// @Tags merchant-branch
// @Accept  json
// @Produce  json
// @Param guid path string true "guid"
// @Success 200 {object} entity.GetMerchantBranchOrdersResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetMerchantBranchOrders(c *gin.Context) {
	response, err := h.storage.MerchantBranchRepo().GetMerchantBranchOrders(&entity.GetMerchantBranchOrdersRequest{
		BranchId: c.Param("guid"),
	})

	if h.handleError(c, err, "Error get merchant-branch orders") {
		return
	}

	c.JSON(200, response)
}
