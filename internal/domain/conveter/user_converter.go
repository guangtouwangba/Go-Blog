package conveter

import (
	"Go-Blog/internal/domain/do"
	"Go-Blog/internal/domain/dto/request"
	"Go-Blog/internal/domain/entity"
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

func (u *UserConverter) EntityToPo(user *entity.User) *do.User {
	return &do.User{
		Email:    user.Email,
		UserName: user.UserName,
		Password: user.Password,
		PhoneNum: user.PhoneNum,
		Describe: user.Describe,
	}
}

func (u *UserConverter) RequestToPo(user *request.UserRegisterRequest) *do.User {
	return &do.User{
		Email:    user.Email,
		UserName: user.UserName,
		Password: user.Password,
		PhoneNum: user.PhoneNum,
		Describe: user.Describe,
	}
}

func (u *UserConverter) PoToEntity(user *do.User) *entity.User {
	return &entity.User{
		Email:    user.Email,
		UserName: user.UserName,
		Password: user.Password,
		PhoneNum: user.PhoneNum,
		Describe: user.Describe,
	}
}
