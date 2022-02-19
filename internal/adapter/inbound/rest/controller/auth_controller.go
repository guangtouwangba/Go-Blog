package controller

import (
	"Go-Blog/internal/app/usecase"
	"Go-Blog/internal/domain/dto/response"
	"Go-Blog/internal/domain/po"
	"Go-Blog/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"time"
)

var jwtSecret = []byte("go-blog")

type AuthController struct {
	BaseController
}

func (a *AuthController) GetToken(c *gin.Context) {
	username := c.Param("name")
	log.Println(username)
	usecases := usecase.UserUseCase{
		UserRepository: &service.UserService{},
	}
	user, err := usecases.GetUserByUserName(username)
	if err != nil {
		response.Error(http.StatusInternalServerError, c)
		return
	}

	log.Println(user)
}

func (a *AuthController) generateToken(username string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := po.UserClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "go-blog",
		},
		TokenType: "level1",
		UserName:  username,
		Password:  password,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*po.UserClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &po.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*po.UserClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
