package routing

import (
	"Go-Blog/internal/adapter/inbound/rest/controller"
	"Go-Blog/internal/app/usecase"
	"Go-Blog/internal/middleware"
	"Go-Blog/internal/service"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	baseController  = controller.BaseController{}
	loginController = controller.UserController{
		UserUseCase: usecase.UserUseCase{
			UserRepository: &service.UserService{},
		},
	}
	articleController = controller.ArticleController{
		ArticleUseCase: usecase.ArticleUseCase{
			ArticleRepository: &service.ArticleService{},
		},
	}
)

func InitRouter() *gin.Engine {
	var router = gin.Default()
	router.Use(middleware.Recover)

	log.Println("Router initialized")
	router.GET("/health", baseController.Health)
	router.POST("/api/v1/login", loginController.UserLogin)
	router.GET("/api/v1/user/:email", loginController.UserInfo)
	router.POST("/api/v1/register", loginController.UserRegister)
	router.GET("/api/v1/articles", articleController.AllArticles)
	return router
}
