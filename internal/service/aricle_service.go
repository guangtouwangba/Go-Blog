package service

import (
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/do"
	"Go-Blog/internal/domain/do/repository"
	"Go-Blog/internal/domain/dto/request"
	"log"
)

type ArticleService struct {
	ArticleRepo repository.ArticleRepository
}

func (a *ArticleService) GetArticle(id int) (*do.ArticleDO, error) {
	return a.ArticleRepo.GetArticle(id)
}

func (a *ArticleService) GetArticles(request *request.ArticleRequest) ([]*do.ArticleDO, error) {
	articlePo := &do.ArticleDO{
		Keywords: request.Keyword,
		State:    request.State,
	}
	pageNum := request.PageNum
	pageSize := request.PageSize
	if request.Archive == 1 {
		pageSize = 1000
	}
	articles := make([]*do.ArticleDO, 0)
	constant.MysqlConnect.Where(articlePo).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&articles)
	log.Println(articles)
	return articles, nil
}
