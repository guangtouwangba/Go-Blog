package routing

type RouterGroup struct {
	UserRouter
	ArticleRouter
	NoticeRouter
	AuthRouter
}

var RouterGroupInstance = new(RouterGroup)
