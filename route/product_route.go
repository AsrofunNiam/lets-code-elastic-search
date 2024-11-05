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

	creditNoteService := service.NewProductService(
		repository.NewProductRepository(),
		db,
		validate,
		elasticClient,
	)
	creditNoteController := controller.NewProductController(creditNoteService)

	router.POST("/elastic/products/process/", creditNoteController.Create)
	router.POST("/elastic/products/sync/", creditNoteController.Sync)

}
