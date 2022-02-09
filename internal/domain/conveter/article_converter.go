package conveter

import (
	"Go-Blog/internal/domain/entity"
	"Go-Blog/internal/domain/po"
)

type ArticleConverter struct {
}

func (a *ArticleConverter) ToArticleEntity() *entity.Article {
	return &entity.Article{}
}

func (a *ArticleConverter) PoToArticleEntity(articlePo *po.Article) *entity.Article {
	return &entity.Article{
		Title:   articlePo.Title,
		Content: articlePo.Content,
	}
}

func (a *ArticleConverter) PosToArticleEntities(articlePos []*po.Article) []*entity.Article {
	var articles []*entity.Article
	for _, articlePo := range articlePos {
		articles = append(articles, a.PoToArticleEntity(articlePo))
	}
	return articles
}
