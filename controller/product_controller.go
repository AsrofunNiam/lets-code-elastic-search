package controller

import (
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	Create(context *gin.Context)
	Sync(context *gin.Context)
}
