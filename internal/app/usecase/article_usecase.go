package usecase

import (
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/entity"
	"Go-Blog/internal/domain/po"
)

type ArticleUseCase struct {
	ArticleRepository po.ArticleRepository
}

func (a *ArticleUseCase) GetArticle(id int) (*po.Article, error) {
	return a.ArticleRepository.GetArticle(id)
}

func (a *ArticleUseCase) GetArticles() ([]*entity.Article, error) {
	articles, err := a.ArticleRepository.GetArticles()
	if err != nil {
		return nil, err
	}

	articleEntities := constant.ArticleConverter.PosToArticleEntities(articles)

	return articleEntities, nil
}
