package routing

import (
	"Go-Blog/internal/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

var ()

func InitRouter() *gin.Engine {
	var Router = gin.Default()
	Router.Use(middleware.Recover)

	log.Println("Router initialized")
	PrivateGroup := Router.Group("/api/v1")

	{
		RouterGroupInstance.InitUserRouter(PrivateGroup)
		RouterGroupInstance.InitArticleRouter(PrivateGroup)
	}

	Router.GET("/health", baseController.Health)
	return Router
}
