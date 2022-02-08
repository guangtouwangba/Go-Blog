package po

import "Go-Blog/internal/domain/entity"

type IUser interface {
	Create(user *User) error
	GetByEmail(email string) (*entity.User, error)
}
