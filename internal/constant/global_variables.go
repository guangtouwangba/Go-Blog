package constant

import (
	"Go-Blog/internal/adapter/outbound/db"
	"Go-Blog/internal/domain/conveter"
)

var (
	UserConverter    = &conveter.UserConverter{}
	ArticleConverter = &conveter.ArticleConverter{}
	tool             = &db.MysqlTool{}
	Connect          = tool.GetInstance().GetDb()
)

const (
	RecordNotExist   = "record not exist"
	LoginFailed      = "login failed"
	GetArticleFailed = "get article failed"
	InvalidParams    = "invalid params"
	GetUserFailed    = "get user failed"
)
