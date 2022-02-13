package routing

import (
	"Go-Blog/internal/adapter/inbound/rest/controller"
	"github.com/gin-gonic/gin"
)

type NoticeRouter struct {
}

var (
	noticeController = controller.NoticeController{}
)

func (s *NoticeRouter) InitNoticeRouter(Router *gin.RouterGroup) {
	noticeRouter := Router.Group("")
	{
		noticeRouter.GET("/notices", noticeController.GetNoticesList)
	}
}
