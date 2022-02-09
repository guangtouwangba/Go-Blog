package service

import (
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/dto/request"
	"Go-Blog/internal/domain/po"
	"log"
)

type ArticleService struct {
	ArticleRepo po.ArticleRepository
}

func (a *ArticleService) GetArticle(id int) (*po.Article, error) {
	return a.ArticleRepo.GetArticle(id)
}

func (a *ArticleService) GetArticles(request *request.ArticleRequest) ([]*po.Article, error) {
	articlePo := &po.Article{
		Keywords: request.Keyword,
		State:    request.State,
	}
	pageNum := request.PageNum
	pageSize := request.PageSize
	if request.Archive == 1 {
		pageSize = 1000
	}
	articles := make([]*po.Article, 0)
	constant.Connect.Where(articlePo).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&articles)
	log.Println(articles)
	return articles, nil
}
