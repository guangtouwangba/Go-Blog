package controller

import (
	"Go-Blog/internal/app/usecase"
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/dto/request"
	"Go-Blog/internal/domain/dto/response"
	"Go-Blog/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserController struct {
	BaseController
}

func (u *UserController) UserLogin(c *gin.Context) {
	req := request.UserLoginRequest{}
	c.BindJSON(&req)
	userUseCase := usecase.UserUseCase{}
	user, err := userUseCase.Login(&req)
	res := response.Response{}
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, res.Error(
			http.StatusInternalServerError,
			constant.LOGIN_FAILED))
		return
	}
	log.Println(req)
	c.JSON(http.StatusOK, res.SuccessWithData(user))
}

func (u *UserController) UserRegister(c *gin.Context) {
	register := request.UserRegisterRequest{}
	c.BindJSON(&register)
	log.Println("request from front end", c.Request.Body)
	userUseCase := usecase.UserUseCase{}
	res, err := userUseCase.Register(&register)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	r := &response.Response{}
	c.JSON(200, r.SuccessWithData(res))

}

func (u *UserController) UserInfo(c *gin.Context) {
	userService := service.UserService{}
	userpo, err := userService.GetByEmail(c.Param("email"))
	if err != nil {
		log.Panicln(err)
	}
	r := &response.Response{}
	c.JSON(200, r.SuccessWithData(userpo))
}
