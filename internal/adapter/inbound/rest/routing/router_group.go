package routing

type RouterGroup struct {
	UserRouter
	ArticleRouter
	NoticeRouter
}

var RouterGroupInstance = new(RouterGroup)
