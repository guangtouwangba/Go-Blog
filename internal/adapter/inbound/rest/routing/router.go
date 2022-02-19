package routing

import (
	"Go-Blog/internal/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter() *gin.Engine {
	var Router = gin.Default()
	Router.Use(middleware.Recover).Use(middleware.RequestLogger)

	log.Println("Router initialized")
	PrivateGroup := Router.Group("/api")

	{
		RouterGroupInstance.InitUserRouter(PrivateGroup)
		RouterGroupInstance.InitArticleRouter(PrivateGroup)
		RouterGroupInstance.InitNoticeRouter(PrivateGroup)
		RouterGroupInstance.InitAuthRoute(PrivateGroup)
	}

	Router.GET("/health", baseController.Health)
	return Router
}
