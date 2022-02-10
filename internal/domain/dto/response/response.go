package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const SUCCESS = 0
const INVALID_PARAMS = 1
const ERROR = 7

type Response struct {
	Code    int64       `json:"code" default:"0"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Result(code int64, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
	return
}

func Success(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)

}

func SuccessWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func Error(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func ErrorWithMsg(msg string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, msg, c)
}

func InvalidParam(c *gin.Context) {
	Result(INVALID_PARAMS, map[string]interface{}{}, "参数错误", c)
}

func InvalidParamWithMsg(msg string, c *gin.Context) {
	Result(INVALID_PARAMS, map[string]interface{}{}, msg, c)
}
