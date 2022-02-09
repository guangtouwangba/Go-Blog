package po

import "Go-Blog/internal/domain/dto/request"

type ArticleRepository interface {
	GetArticle(id int) (*Article, error)
	GetArticles(request *request.ArticleRequest) ([]*Article, error)
}
