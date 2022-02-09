package po

import "Go-Blog/internal/domain/entity"

type UserRepository interface {
	Create(user *User) error
	GetByEmail(email string) (*entity.User, error)
}
