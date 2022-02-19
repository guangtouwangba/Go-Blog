package usecase

import (
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/dto/request"
	"Go-Blog/internal/domain/entity"
	"Go-Blog/internal/domain/po"
	"Go-Blog/internal/service"
	uuid "github.com/satori/go.uuid"
	"log"
)

type UserUseCase struct {
	UserRepository po.UserRepository
}

var JWTService = service.JWTService{}

func (l *UserUseCase) Login(login *request.UserLoginRequest) (*entity.User, error) {
	user, err := l.UserRepository.GetByEmail(login.Email)
	if err != nil || user.Password != login.Password {
		log.Panicln(constant.LoginFailed)
		return nil, err
	}
	return constant.UserConverter.PoToEntity(user), nil
}

// 管理端登陆 可通过用户名或者邮箱
func (l *UserUseCase) AdminLogin(login *request.AdminLoginRequest) (*entity.User, error) {
	userId, err := l.UserRepository.GetUserIdByUsername(login.Account)
	if userId == uuid.Nil {
		userId, err = l.UserRepository.GetUserIdByEmail(login.Account)
	}
	if userId == uuid.Nil {
		// 用户不存在
		log.Panicln(constant.LoginErrorNoRecord.Message)
	}

	if err != nil {
		log.Panicln(constant.LoginErrorNoRecord.Message)
	}

	user, err := l.UserRepository.GetUserById(userId)
	if err != nil {
		// 用户不存在
		log.Panicln(constant.LoginErrorNoRecord.Message)
	}

	if user.Password != login.Password {
		log.Panicln(constant.LoginErrorInvalidParams.Message)
	}

	token, err := JWTService.CheckUserStatus(user)
	if err != nil {
		log.Panicln(constant.LoginErrorTokenExpired.Message)
	}

	_, err = JWTService.ParseToken(token)
	if err != nil {
		log.Panicln(constant.LoginErrorTokenExpired.Message)
	}

	return constant.UserConverter.PoToEntity(user), nil
}

func (l *UserUseCase) Register(register *request.UserRegisterRequest) (*entity.User, error) {
	userpo := constant.UserConverter.RequestToPo(register)
	log.Println(&register)
	err := l.UserRepository.Create(userpo)
	if err != nil {
		log.Panicln(err)
	}

	return constant.UserConverter.PoToEntity(userpo), nil
}

func (l *UserUseCase) GetUserById(id uuid.UUID) (*entity.User, error) {
	user, err := l.UserRepository.GetUserById(id)
	if err != nil {
		log.Panicln(err)
	}
	return constant.UserConverter.PoToEntity(user), nil
}

func (l *UserUseCase) GetUserByUserName(username string) (*entity.User, error) {
	user, err := l.UserRepository.GetUserByUserName(username)
	if err != nil {
		log.Panicln(err)
	}
	return constant.UserConverter.PoToEntity(user), nil
}
