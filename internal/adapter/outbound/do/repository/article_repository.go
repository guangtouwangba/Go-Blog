package repository

import (
	"Go-Blog/internal/adapter/inbound/rest/controller/dto/request"
	"Go-Blog/internal/adapter/outbound/do"
)

type ArticleRepository interface {
	GetArticle(id int) (*do.ArticleDO, error)
	GetArticles(request *request.ArticleRequest) ([]*do.ArticleDO, error)
}
