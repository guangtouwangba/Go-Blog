package service

import (
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/do"
	"Go-Blog/internal/domain/do/repository"
	uuid "github.com/satori/go.uuid"
	"log"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (u *UserService) GetByEmail(email string) (*do.User, error) {
	user := &do.User{}
	res := constant.MysqlConnect.First(user, "email = ?", email)
	if res.Error != nil {
		log.Panicln(constant.LoginErrorEmptyParams.Message)
	}
	//entityUser := constant.UserConverter.PoToEntity(user)
	log.Println(user)
	return user, nil
}

func (u *UserService) GetUserByUserName(username string) (*do.User, error) {
	user := &do.User{}
	res := constant.MysqlConnect.Find(user, "user_name = ?", username)
	log.Println("find user by user name : ", user)
	if res.Error != nil {
		panic(constant.RecordNotExist)
		return nil, res.Error
	}
	return user, nil
}

func (u *UserService) GetUserIdByEmail(email string) (uuid.UUID, error) {
	user, err := u.GetByEmail(email)
	if err != nil {
		log.Panicln(constant.LoginErrorEmptyParams.Message)
		return uuid.Nil, err
	}
	return user.Uuid, nil
}

func (u *UserService) GetUserIdByUsername(username string) (uuid.UUID, error) {
	user, err := u.GetUserByUserName(username)
	if err != nil {
		log.Panicln(constant.LoginErrorEmptyParams.Message)
	}
	return user.Uuid, nil
}

func (u *UserService) GetUserById(id uuid.UUID) (*do.User, error) {
	user := &do.User{}
	res := constant.MysqlConnect.First(user, "uuid = ?", id)
	if res.Error != nil {
		log.Panicln(constant.LoginErrorEmptyParams.Message)
		// return nil, res.Error
	}
	return user, nil
}

func (u *UserService) Create(user *do.User) error {
	res := constant.MysqlConnect.Omit("Id").FirstOrCreate(user, &do.User{Email: user.Email})
	log.Println(user)
	return res.Error
}

func (u *UserService) GetUserToken(user *do.User) (string, error) {
	// TODO: Implement this
	panic("implement me")
	return "", nil
}
