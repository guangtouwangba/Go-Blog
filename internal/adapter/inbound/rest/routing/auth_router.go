package routing

import (
	"Go-Blog/internal/adapter/inbound/rest/controller"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

var AuthController = controller.AuthController{
	BaseController: controller.BaseController{},
}

func (a *AuthRouter) InitAuthRoute(Router *gin.RouterGroup) {
	authRouter := Router.Group("")
	{
		authRouter.GET("/auth/", AuthController.GetToken)
	}
}
