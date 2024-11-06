package route

import (
	"github.com/AsrofunNiam/lets-code-elastic-search/controller"
	"github.com/AsrofunNiam/lets-code-elastic-search/repository"
	"github.com/AsrofunNiam/lets-code-elastic-search/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

func ProductRoute(router *gin.Engine, db *gorm.DB, elasticClient *elastic.Client, validate *validator.Validate) {

	productService := service.NewProductService(
		repository.NewProductRepository(),
		db,
		validate,
		elasticClient,
	)
	productController := controller.NewProductController(productService)

	router.GET("/elastic/products", productController.FindAll)
	router.POST("/elastic/products/process", productController.Create)
	router.POST("/elastic/products/sync", productController.Sync)

}
