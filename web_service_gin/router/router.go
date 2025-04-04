package router

import (
	"example/web-service-gin/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/users", handler.GetUsers)
	}
}
