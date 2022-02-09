package conveter

import (
	"Go-Blog/internal/domain/entity"
)

type ArticleConverter struct {
}

func (a *ArticleConverter) ToArticleEntity() *entity.Article {
	return &entity.Article{}
}

func (a *ArticleConverter) ToArticleEntityList() *[]entity.Article {
	return &[]entity.Article{}
}
