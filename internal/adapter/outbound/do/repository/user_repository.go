package repository

import (
	"Go-Blog/internal/adapter/outbound/do"
	uuid "github.com/satori/go.uuid"
)

type UserRepository interface {
	Create(user *do.User) error
	GetUserIdByEmail(email string) (uuid.UUID, error)
	GetUserByUserName(username string) (*do.User, error)
	GetUserIdByUsername(username string) (uuid.UUID, error)
	GetUserById(id uuid.UUID) (*do.User, error)
	GetByEmail(email string) (*do.User, error)
	GetUserToken(user *do.User) (string, error)
}
