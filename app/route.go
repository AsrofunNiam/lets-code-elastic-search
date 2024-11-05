package app

import (
	route "github.com/AsrofunNiam/lets-code-elastic-search/route"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

func NewRouter(elasticClient *elastic.Client, db *gorm.DB, validate *validator.Validate) *gin.Engine {

	router := gin.New()
	router.UseRawPath = true

	route.ProductRoute(router, db, elasticClient, validate)

	return router
}
