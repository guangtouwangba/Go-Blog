package controller

import (
	"Go-Blog/internal/app/usecase"
	"Go-Blog/internal/domain/dto/response"
	"github.com/gin-gonic/gin"
	"log"
)

type ArticleController struct {
	BaseController
	ArticleUseCase usecase.ArticleUseCase
}

func (a *ArticleController) AllArticles(c *gin.Context) {
	articles, err := a.ArticleUseCase.GetArticles()
	if err != nil {
		log.Println(err)
		response.Error(c)
	}
	log.Println(articles)
	response.SuccessWithData(articles, c)
}
