package usecase

import (
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/dto/request"
	"Go-Blog/internal/domain/entity"
	"Go-Blog/internal/domain/po"
	"log"
)

type UserUseCase struct {
	UserRepository po.UserRepository
}

func (l *UserUseCase) Login(login *request.UserLoginRequest) (*entity.User, error) {
	user, err := l.UserRepository.GetByEmail(login.Email)
	if err != nil || user.Password != login.Password {
		log.Panicln(constant.LoginFailed)
		return nil, err
	}
	return constant.UserConverter.PoToEntity(user), nil
}

func (l *UserUseCase) AdminLogin(login *request.AdminLoginRequest) (*entity.User, error) {
	userId, err := l.UserRepository.GetUserIdByUsername(login.UserName)
	if err != nil {
		log.Panicln(constant.LoginFailed)
		return nil, err
	}
	user, err := l.UserRepository.GetUserById(userId)
	if err != nil || user.Password != login.Password {
		log.Panicln(constant.LoginFailed)
		return nil, err
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

func (l *UserUseCase) GetByEmail(email string) (*entity.User, error) {
	user, err := l.UserRepository.GetByEmail(email)
	if err != nil {
		log.Panicln(err)
	}
	return constant.UserConverter.PoToEntity(user), nil
}
