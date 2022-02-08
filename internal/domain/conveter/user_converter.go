package conveter

import (
	"Go-Blog/internal/domain/dto/request"
	"Go-Blog/internal/domain/entity"
	"Go-Blog/internal/domain/po"
)

type UserConverter struct {
}

func (u *UserConverter) ToEntity(user *request.Register) *entity.User {
	return &entity.User{
		Email:    user.Email,
		UserName: user.UserName,
		Password: user.Password,
		PhoneNum: user.PhoneNum,
		Describe: user.Describe,
	}
}

func (u *UserConverter) EntityToPo(user *entity.User) *po.User {
	return &po.User{
		Email:    user.Email,
		UserName: user.UserName,
		Password: user.Password,
		PhoneNum: user.PhoneNum,
		Describe: user.Describe,
	}
}

func (u *UserConverter) RequestToPo(user *request.Register) *po.User {
	return &po.User{
		Email:    user.Email,
		UserName: user.UserName,
		Password: user.Password,
		PhoneNum: user.PhoneNum,
		Describe: user.Describe,
	}
}
