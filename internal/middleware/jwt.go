package middleware

import (
	"Go-Blog/internal/domain/dto/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TODO: 作为中间件去校验JWT
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.ErrorWithMsg(http.StatusUnauthorized, "未登录或非法访问", c)
			c.Abort()
			return
		}
	}
}
