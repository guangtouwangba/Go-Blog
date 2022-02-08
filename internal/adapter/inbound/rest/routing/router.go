package routing

import (
	"Go-Blog/internal/adapter/inbound/rest/controller"
	"Go-Blog/internal/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	baseController  = controller.BaseController{}
	loginController = controller.UserController{}
)

func InitRouter() *gin.Engine {
	var router = gin.Default()
	router.Use(middleware.Recover)

	log.Println("Router initialized")
	router.GET("/health", baseController.Health)
	router.POST("/api/v1/login", loginController.UserLogin)
	router.GET("/api/v1/user/:email", loginController.UserInfo)
	router.POST("/api/v1/register", loginController.UserRegister)
	return router
}
