package usecase

import (
	"Go-Blog/internal/adapter/inbound/rest/controller/dto/request"
	"Go-Blog/internal/adapter/outbound/do"
	"Go-Blog/internal/adapter/outbound/do/repository"
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/entity"
)

type ArticleUseCase struct {
	ArticleRepository repository.ArticleRepository
}

func (a *ArticleUseCase) GetArticle(id int) (*do.ArticleDO, error) {
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
