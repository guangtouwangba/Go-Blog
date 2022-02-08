package service

import (
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/entity"
	"Go-Blog/internal/domain/po"
)

type UserService struct {
}

func (u *UserService) GetByEmail(email string) (*entity.User, error) {
	return &entity.User{
		Email:    email,
		Password: "123456",
	}, nil
}

func (u *UserService) Create(user *po.User) error {
	res := constant.Connect.Omit("Id").Create(user)
	return res.Error
}
