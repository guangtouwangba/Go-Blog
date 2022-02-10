package conveter

import (
	"Go-Blog/internal/domain/dto/request"
	"Go-Blog/internal/domain/entity"
	"Go-Blog/internal/domain/po"
)

// UserConverter TODO: 用户密码加密解密
type UserConverter struct {
}

func (u *UserConverter) ToEntity(user *request.UserRegisterRequest) *entity.User {
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

func (u *UserConverter) RequestToPo(user *request.UserRegisterRequest) *po.User {
	return &po.User{
		Email:    user.Email,
		UserName: user.UserName,
		Password: user.Password,
		PhoneNum: user.PhoneNum,
		Describe: user.Describe,
	}
}

func (u *UserConverter) PoToEntity(user *po.User) *entity.User {
	return &entity.User{
		Email:    user.Email,
		UserName: user.UserName,
		Password: user.Password,
		PhoneNum: user.PhoneNum,
		Describe: user.Describe,
	}
}
