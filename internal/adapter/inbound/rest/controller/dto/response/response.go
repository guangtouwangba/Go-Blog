package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const SUCCESS = 0
const InvalidParams = 1
const ERROR = 7

type Response struct {
	Code    int64       `json:"code" default:"0"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Result(httpStatus int, code int64, data interface{}, msg string, c *gin.Context) {
	c.JSON(httpStatus, Response{
		code,
		msg,
		data,
	})
	return
}

func Success(c *gin.Context) {
	Result(http.StatusOK, SUCCESS, map[string]interface{}{}, "操作成功", c)

}

func SuccessWithData(data interface{}, c *gin.Context) {
	Result(http.StatusOK, SUCCESS, data, "操作成功", c)
}

func Error(httpStatus int, c *gin.Context) {
	Result(httpStatus, ERROR, map[string]interface{}{}, "操作失败", c)
}

func ErrorWithMsg(httpStatus int, msg string, c *gin.Context) {
	Result(httpStatus, ERROR, map[string]interface{}{}, msg, c)
}

func InvalidParam(httpStatus int, c *gin.Context) {
	Result(httpStatus, InvalidParams, map[string]interface{}{}, "参数错误", c)
}

func InvalidParamWithMsg(httpStatus int, msg string, c *gin.Context) {
	Result(httpStatus, InvalidParams, map[string]interface{}{}, msg, c)
}

// 登陆失败
//func LoginFail(msg string, c *gin.Context) {
//	Result(INVALID_PARAMS, map[string]interface{}{}, msg, c)
//}
