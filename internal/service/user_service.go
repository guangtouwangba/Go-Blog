package service

import (
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/entity"
	"Go-Blog/internal/domain/po"
	"log"
)

type UserService struct {
}

func (u *UserService) GetByEmail(email string) (*entity.User, error) {
	user := &po.User{}
	res := constant.Connect.First(user, "email = ?", email)
	if res.Error != nil {
		panic(constant.RECORD_NOT_EXIST)
		return nil, res.Error
	}
	log.Println(user)
	entityUser := constant.UserConverter.PoToEntity(user)
	log.Println(entityUser)
	return entityUser, nil
}

func (u *UserService) Create(user *po.User) error {
	res := constant.Connect.Omit("Id").FirstOrCreate(user)
	return res.Error
}
