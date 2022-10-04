package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tfpolachini/go-crud-example/domain/service"
)

type ProductController struct {
	service service.ProductServiceInterface
}

func NewProductController(service service.ProductServiceInterface) *ProductController {
	return &ProductController{service}
}

func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	var inputDto service.CreateProductInputDto

	err := c.BindJSON(&inputDto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	outputDto, err := ctrl.service.CreateProduct(inputDto)
	if err != nil {
		c.Status(HttpStatusOf(err))
		return
	}

	c.JSON(http.StatusCreated, outputDto)
}
