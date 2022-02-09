package service

import (
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/po"
	"log"
)

type ArticleService struct {
	ArticleRepo po.ArticleRepository
}

func (a *ArticleService) GetArticle(id int) (*po.Article, error) {
	return a.ArticleRepo.GetArticle(id)
}

func (a *ArticleService) GetArticles() ([]*po.Article, error) {
	articles := make([]*po.Article, 0)
	constant.Connect.Find(&articles)
	log.Println(articles)
	return articles, nil
}
