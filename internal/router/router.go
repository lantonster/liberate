package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lantonster/liberate/internal/handler"
	"github.com/lantonster/liberate/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title						Liberate API
//	@description				Liberate API for internal services
//	@version					1.0
//	@host						localhost:8080
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//	@securityDefinitions.apikey	apiKey
//	@in							header
//	@name						Authorization

func RegisterRoutes(r *gin.Engine, service *service.Service, handler *handler.Handler) {

	r.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Hello, World!"}) })
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/users/register", handler.UserHandler.Register)
}
