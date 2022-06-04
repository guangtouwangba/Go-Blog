package conveter

import (
	"Go-Blog/internal/domain/do"
	"Go-Blog/internal/domain/entity"
)

type ArticleConverter struct {
}

func (a *ArticleConverter) ToArticleEntity() *entity.Article {
	return &entity.Article{}
}

func (a *ArticleConverter) PoToArticleEntity(articlePo *do.ArticleDO) *entity.Article {
	return &entity.Article{
		Title:   articlePo.Title,
		Content: articlePo.Content,
	}
}

func (a *ArticleConverter) PosToArticleEntities(articlePos []*do.ArticleDO) []*entity.Article {
	var articles []*entity.Article
	for _, articlePo := range articlePos {
		articles = append(articles, a.PoToArticleEntity(articlePo))
	}
	return articles
}
