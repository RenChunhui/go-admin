package router

import (
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.New()

	api := r.Group("/api")
	{
		api.POST("/account/login")
	}

	return r
}
