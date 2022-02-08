package usecase

import (
	"Go-Blog/internal/domain/dto/request"
	"Go-Blog/internal/domain/entity"
	"fmt"
	"github.com/wenlng/go-captcha/captcha"
)

type LoginUseCase struct {
}

func (l *LoginUseCase) Login(login *request.Login) (*entity.User, error) {
	capt := captcha.GetCaptcha()
	fmt.Println(capt)
	return new(entity.User), nil
}
