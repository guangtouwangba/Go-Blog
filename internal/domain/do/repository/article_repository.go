package repository

import (
	"Go-Blog/internal/domain/do"
	"Go-Blog/internal/domain/dto/request"
)

type ArticleRepository interface {
	GetArticle(id int) (*do.ArticleDO, error)
	GetArticles(request *request.ArticleRequest) ([]*do.ArticleDO, error)
}
