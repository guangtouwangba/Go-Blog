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
	RecordNotExist   = "该记录不存在"
	LoginFailed      = "登陆失败"
	GetArticleFailed = "获取文章失败"
	InvalidParams    = "无效的请求值"
	GetUserFailed    = "获取用户信息失败"
)

var (
	// 登陆相关异常
	Login0000 = GlobolErrCode{Code: "Login_0000", Message: "请输入账户或密码"}
	Login0001 = GlobolErrCode{Code: "Login_0001", Message: "账号不存在"}
	Login0002 = GlobolErrCode{Code: "Login_0002", Message: "登录密码错误"}
	Login0003 = GlobolErrCode{Code: "Login_0003", Message: "账号已被禁用"}
	Login0004 = GlobolErrCode{Code: "Login_0004", Message: "未知错误"}
)
