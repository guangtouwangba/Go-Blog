package service

import "Go-Blog/internal/domain/po"

type ArticleService struct {
	ArticleRepo po.ArticleRepository
}

func (a *ArticleService) GetArticle(id int) (*po.Article, error) {
	return a.ArticleRepo.GetArticle(id)
}

func (a *ArticleService) GetArticles() ([]*po.Article, error) {
	return a.ArticleRepo.GetArticles()
}
