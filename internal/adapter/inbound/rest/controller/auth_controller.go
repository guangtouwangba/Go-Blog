package controller

import (
	"Go-Blog/internal/service"
)

var jwtService = service.JWTService{}

type AuthController struct {
	BaseController
}

//func (a *AuthController) GetToken(c *gin.Context) {
//	username := c.Param("name")
//	log.Println(username)
//	usecases := usecase.UserUseCase{
//		UserRepository: &service.UserService{},
//	}
//	user, err := usecases.GetUserByUserName(username)
//	if err != nil {
//		response.Error(http.StatusInternalServerError, c)
//		return
//	}
//
//	log.Println(user)
//}
//
//func (a *AuthController) generateToken(username string, password string) (string, error) {
//	return jwtService.GenerateJWTToken(username, password)
//}
