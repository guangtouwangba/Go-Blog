package routing

import (
	"Go-Blog/internal/adapter/inbound/rest/controller"
	"Go-Blog/internal/app/usecase"
	"Go-Blog/internal/service"
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {
}

var (
	articleController = controller.ArticleController{
		ArticleUseCase: usecase.ArticleUseCase{
			ArticleRepository: &service.ArticleService{},
		},
	}
)

func (a *ArticleRouter) InitArticleRouter(router *gin.RouterGroup) {
	articleRouter := router.Group("")
	{
		articleRouter.GET("/articleList", articleController.ArticleList)

	}
}
