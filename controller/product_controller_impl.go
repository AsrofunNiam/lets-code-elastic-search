package controller

import (
	"net/http"

	"github.com/AsrofunNiam/lets-code-elastic-search/helper"
	"github.com/AsrofunNiam/lets-code-elastic-search/model/web"
	"github.com/AsrofunNiam/lets-code-elastic-search/service"
	"github.com/gin-gonic/gin"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Create(c *gin.Context) {
	request := web.ProductCreateRequest{}
	helper.ReadFromRequestBody(c, &request)

	controller.ProductService.Create(&request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Process create product successfully",
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *ProductControllerImpl) Sync(c *gin.Context) {

	controller.ProductService.Sync(c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Process sync product successfully",
	}

	c.JSON(http.StatusOK, webResponse)
}