package controller

import (
	"Go-Blog/internal/app/usecase"
	"Go-Blog/internal/domain/dto/request"
	"Go-Blog/internal/domain/dto/response"
	"github.com/gin-gonic/gin"
	"log"
)

type ArticleController struct {
	BaseController
	ArticleUseCase usecase.ArticleUseCase
}

func (a *ArticleController) AllArticles(c *gin.Context) {
	articleRequest := &request.ArticleRequest{}
	err := c.Bind(articleRequest)
	if err != nil {
		response.ErrorWithMsg("Invalid request", c)
	}
	articles, err := a.ArticleUseCase.GetArticles(articleRequest)
	if err != nil {
		log.Println(err)
		response.Error(c)
	}
	log.Println(articles)
	response.SuccessWithData(articles, c)
}
