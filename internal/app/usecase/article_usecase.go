package usecase

import "Go-Blog/internal/domain/po"

type ArticleUseCase struct {
	ArticleRepository po.ArticleRepository
}
