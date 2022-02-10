package po

import uuid "github.com/satori/go.uuid"

type UserRepository interface {
	Create(user *User) error
	GetUserIdByEmail(email string) (uuid.UUID, error)
	GetUserIdByUsername(username string) (uuid.UUID, error)
	GetUserById(id uuid.UUID) (*User, error)
	GetByEmail(email string) (*User, error)
}
