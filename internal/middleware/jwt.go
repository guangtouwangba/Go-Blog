package middleware

import (
	"Go-Blog/config"
	"Go-Blog/internal/domain/do"
	"Go-Blog/internal/domain/dto/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"time"
)

// TODO: 作为中间件去校验JWT
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.ErrorWithMsg(http.StatusUnauthorized, "未登录或非法访问", c)
			c.Abort()
			return
		}
		log.Println("解析到携带的token为：", token)
		j := NewJWT()
		_, err := j.ParseToken(token)
		if err != nil {
			if err == jwt.ErrSignatureInvalid || err == jwt.ErrInvalidKey {
				response.ErrorWithMsg(http.StatusUnauthorized, "签名无效", c)
				c.Abort()
				return
			}

			response.ErrorWithMsg(http.StatusUnauthorized, "未登录或非法访问", c)
		}

		if token == "" {
			response.ErrorWithMsg(http.StatusUnauthorized, "未登录或非法访问", c)
			c.Abort()
			return
		}
	}
}

type JWT struct {
	SigningKey []byte
}

// NewToken : 生成一个token
func (j *JWT) NewToken(claims *do.UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken : 解析token
func (j *JWT) ParseToken(token string) (*do.UserClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &do.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*do.UserClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// RefreshToken : 刷新token
func (j *JWT) RefreshToken(token string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	tokenClaims, err := jwt.ParseWithClaims(token, &do.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*do.UserClaims); ok && tokenClaims.Valid {
			jwt.TimeFunc = time.Now
			claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
			return j.NewToken(claims)
		}
	}
	return "", err
}

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(config.GetYamlConfig().Secret),
	}
}
