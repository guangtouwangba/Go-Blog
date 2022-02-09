package po

type ArticleRepository interface {
	GetArticle(id int) (*Article, error)
	GetArticles() ([]*Article, error)
}
