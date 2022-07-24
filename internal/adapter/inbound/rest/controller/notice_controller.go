package controller

import (
	"Go-Blog/internal/adapter/inbound/rest/controller/dto/response"
	"github.com/gin-gonic/gin"
)

/**
获取消息相关
*/
type NoticeController struct {
	BaseController
}

// GetNoticesList 获取消息列表
func (u *NoticeController) GetNoticesList(c *gin.Context) {

	response.SuccessWithData(nil, c)
}
