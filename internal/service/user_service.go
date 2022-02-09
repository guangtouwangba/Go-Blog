package service

import (
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/po"
	"log"
)

type UserService struct {
	UserRepository po.UserRepository
}

func (u *UserService) GetByEmail(email string) (*po.User, error) {
	user := &po.User{}
	res := constant.Connect.First(user, "email = ?", email)
	if res.Error != nil {
		panic(constant.RecordNotExist)
		return nil, res.Error
	}
	//entityUser := constant.UserConverter.PoToEntity(user)
	log.Println(user)
	return user, nil
}

func (u *UserService) Create(user *po.User) error {
	res := constant.Connect.Omit("Id").FirstOrCreate(user, &po.User{Email: user.Email})
	log.Println(user)
	return res.Error
}
