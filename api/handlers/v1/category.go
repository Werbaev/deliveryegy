package v1

import (
	"github.com/anmimos/delivery/storage/entity"
	"github.com/gin-gonic/gin"
)

// @Router /v1/category/create [POST]
// @Summary Create category
// @Description API that create category
// @Tags category
// @Accept  json
// @Produce  json
// @Param body body entity.CreateCategoryRequest true "body"
// @Success 200 {object} entity.CreateCategoryResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateCategory(c *gin.Context) {
	data := entity.CreateCategoryRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.CategoryRepo().CreateCategory(&entity.CreateCategoryRequest{
		Name:       data.Name,
		MerchantId: data.MerchantId,
	})

	if h.handleError(c, err, "Error create category") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/category/get/{guid} [GET]
// @Summary Get category
// @Description API that get category
// @Tags category
// @Accept  json
// @Produce  json
// @Param guid path string true "guid"
// @Success 200 {object} entity.GetCategoryResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCategory(c *gin.Context) {
	response, err := h.storage.CategoryRepo().GetCategory(&entity.GetCategoryRequest{
		Guid: c.Param("guid"),
	})

	if h.handleError(c, err, "Error get category") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/category/list [GET]
// @Summary List Categories
// @Description API list Categories
// @Tags category
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.ListCategoryResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ListCategories(c *gin.Context) {
	response, err := h.storage.CategoryRepo().ListCategories(&entity.ListCategoryRequest{})

	if h.handleError(c, err, "Error get list category") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/category/update [PUT]
// @Summary Update category
// @Description API that update category
// @Tags category
// @Accept  json
// @Produce  json
// @Param body body entity.UpdateCategoryRequest true "body"
// @Success 200 {object} entity.UpdateCategoryResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateCategory(c *gin.Context) {
	data := entity.UpdateCategoryRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.CategoryRepo().UpdateCategory(&entity.UpdateCategoryRequest{
		Name:       data.Name,
		Guid:       data.Guid,
		MerchantId: data.MerchantId,
	})

	if h.handleError(c, err, "Error update category") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/category/delete [DELETE]
// @Summary Delete category
// @Description API that delete category
// @Tags category
// @Accept  json
// @Produce  json
// @Param body body entity.DeleteCategoryRequest true "body"
// @Success 200 {object} entity.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteCategory(c *gin.Context) {
	data := entity.DeleteCategoryRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.CategoryRepo().DeleteCategory(&entity.DeleteCategoryRequest{
		Guid: data.Guid,
	})

	if h.handleError(c, err, "Error delete category") {
		return
	}

	c.JSON(200, response)
}
