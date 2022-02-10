package service

import (
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/po"
	uuid "github.com/satori/go.uuid"
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

func (u *UserService) GetUserByUserName(username string) (*po.User, error) {
	user := &po.User{}
	res := constant.Connect.Find(user, "username = ?", username)
	if res.Error != nil {
		panic(constant.RecordNotExist)
		return nil, res.Error
	}
	return user, nil
}

func (u *UserService) GetUserIdByEmail(email string) (uuid.UUID, error) {
	user, err := u.GetByEmail(email)
	if err != nil {
		log.Panicln(constant.GetUserFailed)
	}
	return user.Uuid, nil
}

func (u *UserService) GetUserIdByUsername(username string) (uuid.UUID, error) {
	user, err := u.GetUserByUserName(username)
	if err != nil {
		log.Panicln(constant.GetUserFailed)
	}
	return user.Uuid, nil
}

func (u *UserService) GetUserById(id uuid.UUID) (*po.User, error) {
	user := &po.User{}
	res := constant.Connect.First(user, "uuid = ?", id)
	if res.Error != nil {
		panic(constant.RecordNotExist)
		return nil, res.Error
	}
	return user, nil
}

func (u *UserService) Create(user *po.User) error {
	res := constant.Connect.Omit("Id").FirstOrCreate(user, &po.User{Email: user.Email})
	log.Println(user)
	return res.Error
}
