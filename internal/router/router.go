package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/renchunhui/go-admin/internal/router/api"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	if os.Getenv("GIN_MODE") == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	user := api.NewUser()

	api := r.Group("/api")
	{

		// api.POST("/account/singin")
		// api.POST("/account/signup")
		// api.POST("/account/signout")

		api.POST("/users", user.Create)
		api.GET("/users/:id", user.Get)
		api.PATCH("/users/:id/state", user.Update)
		api.DELETE("/users/:id", user.Delete)

		// api.POST("/roles")
		// api.GET("/roles/:id")
		// api.PATCH("/roles/:id/state")
		// api.DELETE("/roles/:id")

		// api.POST("/permissions")
		// api.GET("/permissions/:id")
		// api.PATCH("/permissions/:id")
		// api.DELETE("/permissions/:id")

		// api.POST("/menus")
		// api.GET("/menus/:id")
		// api.PATCH("/menus/:id")
		// api.DELETE("/menus/:id")

		// api.POST("/tags")
		// api.GET("/tags/:id")
		// api.PATCH("/tags/:id")
		// api.DELETE("/tags/:id")

		// api.POST("/categories")
		// api.GET("/categories/:id")
		// api.PATCH("/categories/:id")
		// api.DELETE("/categories/:id")

		// api.POST("/dictionary")
		// api.GET("/dictionary/:id")
		// api.PATCH("/dictionary/:id")
		// api.DELETE("/dictionary/:id")
	}

	return r
}
