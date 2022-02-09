package controller

import (
	"Go-Blog/internal/app/usecase"
	"Go-Blog/internal/domain/dto/request"
	"Go-Blog/internal/domain/dto/response"
	"github.com/gin-gonic/gin"
	"log"
)

type UserController struct {
	BaseController
	UserUseCase usecase.UserUseCase
}

func (u *UserController) UserLogin(c *gin.Context) {
	req := request.UserLoginRequest{}
	c.BindJSON(&req)
	user, err := u.UserUseCase.Login(&req)
	if err != nil {
		log.Println(err)
		response.Error(c)
		return
	}
	log.Println(req)
	response.SuccessWithData(user, c)
}

func (u *UserController) UserRegister(c *gin.Context) {
	register := request.UserRegisterRequest{}
	c.BindJSON(&register)
	log.Println("request from front end", c.Request.Body)
	res, err := u.UserUseCase.Register(&register)
	if err != nil {
		log.Println(err)
		response.Error(c)
		return
	}
	response.SuccessWithData(res, c)
}

func (u *UserController) UserInfo(c *gin.Context) {
	userpo, err := u.UserUseCase.GetByEmail(c.Param("email"))
	if err != nil {
		response.Error(c)
		log.Panicln(err)
	}
	response.SuccessWithData(userpo, c)
}
