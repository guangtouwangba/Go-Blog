package usecase

import (
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/dto/request"
	"Go-Blog/internal/domain/entity"
	"Go-Blog/internal/domain/po"
)

type ArticleUseCase struct {
	ArticleRepository po.ArticleRepository
}

func (a *ArticleUseCase) GetArticle(id int) (*po.Article, error) {
	return a.ArticleRepository.GetArticle(id)
}

func (a *ArticleUseCase) GetArticles(request *request.ArticleRequest) ([]*entity.Article, error) {
	articles, err := a.ArticleRepository.GetArticles(request)
	if err != nil {
		return nil, err
	}

	articleEntities := constant.ArticleConverter.PosToArticleEntities(articles)
	if len(articleEntities) == 0 {
		return make([]*entity.Article, 0), nil
	}
	return articleEntities, nil
}
