package constant

import (
	"Go-Blog/internal/adapter/outbound/db"
	"Go-Blog/internal/adapter/outbound/redis"
	"Go-Blog/internal/domain/conveter"
)

var (
	UserConverter    = &conveter.UserConverter{}
	ArticleConverter = &conveter.ArticleConverter{}
	tool             = &db.MysqlTool{}
	Connect          = tool.GetInstance().GetDb()
	RedisConnect     = redis.CreateRedisConnect()
)

const (
	RecordNotExist   = "record not exist"
	LoginFailed      = "login failed"
	GetArticleFailed = "get article failed"
	InvalidParams    = "invalid params"
	GetUserFailed    = "get user failed"
)

var (
	// 登陆相关异常
	Login_0000 = GlobolErrCode{Code: "Login_0000", Message: "请输入账户或密码"}
	Login_0001 = GlobolErrCode{Code: "Login_0001", Message: "账号不存在"}
	Login_0002 = GlobolErrCode{Code: "Login_0002", Message: "登录密码错误"}
	Login_0003 = GlobolErrCode{Code: "Login_0003", Message: "该用户无管理员权限"}
	Login_0004 = GlobolErrCode{Code: "Login_0004", Message: "未知错误"}
)
