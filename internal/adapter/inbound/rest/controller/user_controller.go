package controller

import (
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/dto/request"
	"Go-Blog/internal/domain/dto/response"
	"Go-Blog/internal/service"
	"github.com/gin-gonic/gin"
	"log"
)

type UserController struct {
	BaseController
}

func (u *UserController) UserLogin(c *gin.Context) {
	json := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	c.BindJSON(&json)
	log.Println(json)
}

func (u *UserController) UserRegister(c *gin.Context) {
	register := request.Register{}
	c.BindJSON(&register)
	log.Println("request from front end", c.Request.Body)
	userpo := constant.UserConverter.RequestToPo(&register)
	log.Println(&register)
	userService := service.UserService{}
	err := userService.Create(userpo)
	if err != nil {
		log.Panicln(err)
	}
	r := &response.Response{}
	c.JSON(200, r.SuccessWithData(userpo))

}
