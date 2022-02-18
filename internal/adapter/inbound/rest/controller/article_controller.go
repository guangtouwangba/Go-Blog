package controller

import (
	"Go-Blog/internal/app/usecase"
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/dto/request"
	"Go-Blog/internal/domain/dto/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ArticleController struct {
	BaseController
	ArticleUseCase usecase.ArticleUseCase
}

func (a *ArticleController) ArticleList(c *gin.Context) {
	articleRequest := &request.ArticleRequest{}
	err := c.Bind(articleRequest)
	if err != nil {
		response.ErrorWithMsg(http.StatusBadRequest, constant.InvalidParams, c)
	}
	articles, err := a.ArticleUseCase.GetArticles(articleRequest)
	if err != nil {
		log.Println(err)
		response.Error(http.StatusInternalServerError, c)
	}
	response.SuccessWithData(articles, c)
}
