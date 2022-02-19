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
	LoginErrorEmptyParams     = GlobolErrCode{Code: "Login_0000", Message: "请输入账户或密码"}
	LoginErrorNoRecord        = GlobolErrCode{Code: "Login_0001", Message: "账号不存在"}
	LoginErrorInvalidParams   = GlobolErrCode{Code: "Login_0002", Message: "登录密码错误"}
	LoginErrorAccountDisabled = GlobolErrCode{Code: "Login_0003", Message: "账号已被禁用"}
	LoginErrorUnknown         = GlobolErrCode{Code: "Login_0004", Message: "未知错误"}
	LoginErrorTokenExpired    = GlobolErrCode{Code: "Login_0005", Message: "账号登陆已过期，请重新登陆"}
)
