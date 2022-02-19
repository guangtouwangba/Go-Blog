package middleware

import (
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/dto/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			//封装通用json返回
			responseGenerator(errorToString(r), c)
			//response.ErrorWithMsg(http.StatusInternalServerError, errorToString(r), c)
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

func responseGenerator(errMsg string, c *gin.Context) {
	if errMsg == constant.LoginErrorTokenExpired.Message {
		response.ErrorWithMsg(http.StatusUnauthorized, errMsg, c)
	} else {
		response.ErrorWithMsg(http.StatusBadRequest, errMsg, c)
	}

}
