package routing

type RouterGroup struct {
	UserRouter
	ArticleRouter
}

var RouterGroupInstance = new(RouterGroup)
