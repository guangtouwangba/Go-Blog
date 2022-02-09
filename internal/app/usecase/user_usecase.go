package usecase

import (
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/dto/request"
	"Go-Blog/internal/domain/entity"
	"Go-Blog/internal/service"
	"log"
)

type UserUseCase struct {
}

func (l *UserUseCase) Login(login *request.UserLoginRequest) (*entity.User, error) {
	userService := &service.UserService{}
	user, err := userService.GetByEmail(login.Email)
	if err != nil || user.Password != login.Password {
		log.Panicln(constant.LOGIN_FAILED)
		return nil, err
	}
	return user, nil
}

func (l *UserUseCase) Register(register *request.UserRegisterRequest) (*entity.User, error) {
	userpo := constant.UserConverter.RequestToPo(register)
	log.Println(&register)
	userService := service.UserService{}
	err := userService.Create(userpo)
	if err != nil {
		log.Panicln(err)
	}

	return constant.UserConverter.PoToEntity(userpo), nil
}
