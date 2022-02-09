package po

type UserRepository interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
}
