package controller

import (
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	FindAll(context *gin.Context)
	Create(context *gin.Context)
	Sync(context *gin.Context)
}
