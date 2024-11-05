package service

import (
	"github.com/AsrofunNiam/lets-code-elastic-search/model/web"
	"github.com/gin-gonic/gin"
)

type ProductService interface {
	Create(request *web.ProductCreateRequest, c *gin.Context) web.ProductResponse
	Sync(c *gin.Context)
	Delete(id *string, c *gin.Context)
}
