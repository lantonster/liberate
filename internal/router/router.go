package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lantonster/liberate/internal/handler"
	"github.com/lantonster/liberate/internal/service"
)

func RegisterRoutes(r *gin.Engine, service *service.Service, handler *handler.Handler) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.POST("/register", handler.UserHandler.Register)
}
