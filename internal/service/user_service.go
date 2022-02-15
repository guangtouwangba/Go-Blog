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
		log.Panicln(constant.Login_0001.Message)
	}
	//entityUser := constant.UserConverter.PoToEntity(user)
	log.Println(user)
	return user, nil
}

func (u *UserService) GetUserByUserName(username string) (*po.User, error) {
	user := &po.User{}
	res := constant.Connect.Find(user, "user_name = ?", username)
	if res.Error != nil {
		panic(constant.RecordNotExist)
		return nil, res.Error
	}
	return user, nil
}

func (u *UserService) GetUserIdByEmail(email string) (uuid.UUID, error) {
	user, err := u.GetByEmail(email)
	if err != nil {
		log.Panicln(constant.Login_0001.Message)
		return uuid.Nil, err
	}
	return user.Uuid, nil
}

func (u *UserService) GetUserIdByUsername(username string) (uuid.UUID, error) {
	user, err := u.GetUserByUserName(username)
	if err != nil {
		log.Panicln(constant.Login_0001.Message)
	}
	return user.Uuid, nil
}

func (u *UserService) GetUserById(id uuid.UUID) (*po.User, error) {
	user := &po.User{}
	res := constant.Connect.First(user, "uuid = ?", id)
	if res.Error != nil {
		log.Panicln(constant.Login_0001.Message)
		// return nil, res.Error
	}
	return user, nil
}

func (u *UserService) Create(user *po.User) error {
	res := constant.Connect.Omit("Id").FirstOrCreate(user, &po.User{Email: user.Email})
	log.Println(user)
	return res.Error
}
