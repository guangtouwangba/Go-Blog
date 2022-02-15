package controller

import (
	"Go-Blog/internal/app/usecase"
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/dto/request"
	"Go-Blog/internal/domain/dto/response"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
)

type UserController struct {
	BaseController
	UserUseCase usecase.UserUseCase
}

// UserLogin 博客端 普通用户登陆
func (u *UserController) UserLogin(c *gin.Context) {
	req := request.UserLoginRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		response.InvalidParamWithMsg(constant.InvalidParams, c)
	}
	user, err := u.UserUseCase.Login(&req)
	if err != nil {
		log.Println(err)
		response.Error(c)
		return
	}
	log.Println(req)
	response.SuccessWithData(user, c)
}

// AdminLogin 管理端 管理员登陆
func (u *UserController) AdminLogin(c *gin.Context) {
	req := request.AdminLoginRequest{}
	err := c.BindJSON(&req)
	if err != nil || req.Account == "" {
		response.LoginFail(constant.Login_0000.Message, c)
		return
	}
	user, err := u.UserUseCase.AdminLogin(&req)

	if err != nil {
		log.Println(err)
		response.LoginFail(err.Error(), c)
		return
	}
	log.Println(req)
	response.SuccessWithData(user, c)
}

// GetCurrentUser
func (u *UserController) GetCurrentUser(c *gin.Context) {
	// todo 获取当前用户
	response.SuccessWithData(nil, c)
}

func (u *UserController) UserRegister(c *gin.Context) {
	register := request.UserRegisterRequest{}
	err := c.BindJSON(&register)
	if err != nil {
		response.InvalidParamWithMsg(constant.InvalidParams, c)
		log.Panicln(err)
	}
	log.Println("request from front end", c.Request.Body)
	res, err := u.UserUseCase.Register(&register)
	if err != nil {
		log.Println(err)
		response.Error(c)
		return
	}
	response.SuccessWithData(res, c)
}

// UserInfoById TODO: 后期删除该方法，通过UUID获取用户信息
func (u *UserController) UserInfoById(c *gin.Context) {
	userpo, err := u.UserUseCase.GetUserById(uuid.FromStringOrNil(c.Param("id")))
	if err != nil {
		response.Error(c)
		log.Panicln(err)
	}
	response.SuccessWithData(userpo, c)
}
