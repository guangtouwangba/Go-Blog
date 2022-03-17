package controller

import (
	"Go-Blog/internal/app/usecase"
	"Go-Blog/internal/domain/dto/response"
	"Go-Blog/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var jwtService = service.JWTService{}

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
