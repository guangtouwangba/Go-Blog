package routing

import (
	"Go-Blog/internal/adapter/inbound/rest/controller"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	baseController  = controller.BaseController{}
	loginController = controller.UserController{}
)

func InitRouter() *gin.Engine {
	var router = gin.Default()
	log.Println("Router initialized")
	router.GET("/health", baseController.Health)
	router.POST("/api/v1/login", loginController.UserLogin)
	router.POST("/api/v1/register", loginController.UserRegister)
	return router
}
