package v1

import (
	"github.com/anmimos/delivery/storage/entity"
	"github.com/gin-gonic/gin"
)

// @Router /v1/product/create [POST]
// @Summary Create product
// @Description API that create product
// @Tags product
// @Accept  json
// @Produce  json
// @Param body body entity.CreateProductRequest true "body"
// @Success 200 {object} entity.CreateProductResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateProduct(c *gin.Context) {
	data := entity.CreateProductRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.ProductRepo().CreateProduct(&entity.CreateProductRequest{
		Name:       data.Name,
		Price:      data.Price,
		CategoryId: data.CategoryId,
		Option:     data.Option,
		Image:      data.Image,
	})

	if h.handleError(c, err, "Error create product") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/product/get/{guid} [GET]
// @Summary Get product
// @Description API that get product
// @Tags product
// @Accept  json
// @Produce  json
// @Param guid path string true "guid"
// @Success 200 {object} entity.GetProductResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetProduct(c *gin.Context) {
	response, err := h.storage.ProductRepo().GetProduct(&entity.GetProductRequest{
		Guid: c.Param("guid"),
	})

	if h.handleError(c, err, "Error get product") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/product/list [GET]
// @Summary List Products
// @Description API list products
// @Tags product
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.ListProductResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ListProducts(c *gin.Context) {
	response, err := h.storage.ProductRepo().ListProducts(&entity.ListProductRequest{})

	if h.handleError(c, err, "Error get list products") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/product/update [PUT]
// @Summary Update product
// @Description API that update product
// @Tags product
// @Accept  json
// @Produce  json
// @Param body body entity.UpdateProductRequest true "body"
// @Success 200 {object} entity.UpdateProductResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateProduct(c *gin.Context) {
	data := entity.UpdateProductRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.ProductRepo().UpdateProduct(&entity.UpdateProductRequest{
		Name:       data.Name,
		Guid:       data.Guid,
		Price:      data.Price,
		CategoryId: data.CategoryId,
		Option:     data.Option,
		Image:      data.Image,
	})

	if h.handleError(c, err, "Error update product") {
		return
	}

	c.JSON(200, response)
}

// @Router /v1/product/delete [DELETE]
// @Summary Delete product
// @Description API that delete product
// @Tags product
// @Accept  json
// @Produce  json
// @Param body body entity.DeleteProductRequest true "body"
// @Success 200 {object} entity.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteProduct(c *gin.Context) {
	data := entity.DeleteProductRequest{}

	err := c.ShouldBindJSON(&data)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	response, err := h.storage.ProductRepo().DeleteProduct(&entity.DeleteProductRequest{
		Guid: data.Guid,
	})

	if h.handleError(c, err, "Error delete product") {
		return
	}

	c.JSON(200, response)
}
