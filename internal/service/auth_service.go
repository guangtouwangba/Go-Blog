package service

import (
	"Go-Blog/config"
	"Go-Blog/internal/constant"
	"Go-Blog/internal/domain/po"
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

type JWTService struct {
}

func (j *JWTService) GetKeyFromRedis(userId uuid.UUID) (string, error) {
	jwt, err := constant.RedisConnect.Get(userId.String()).Result()
	return jwt, err
}

func (j *JWTService) SetKeyToRedis(userId uuid.UUID, jwt string, duration time.Duration) (string, error) {
	value, err := constant.RedisConnect.Set(userId.String(), jwt, duration*time.Second).Result()
	return value, err
}

func (j *JWTService) GenerateJWTToken(username string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(config.GetYamlConfig().Expire) * time.Second)
	claims := po.UserClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "go-blog",
		},
		TokenType: "level1",
		UserName:  username,
		Password:  password,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Println(config.GetYamlConfig().JWTConfig.Secret)
	token, err := tokenClaims.SignedString([]byte(config.GetYamlConfig().JWTConfig.Secret))
	log.Println(token)
	return token, err
}

func (j *JWTService) ParseToken(token string) (*po.UserClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &po.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetYamlConfig().JWTConfig.Secret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*po.UserClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			log.Panicln(constant.LoginErrorUnknown.Message)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			log.Panicln(constant.LoginErrorTokenExpired.Message)
		} else {
			log.Panicln(constant.LoginErrorUnknown.Message)
		}
	} else {
		log.Panicln(constant.LoginErrorUnknown.Message)
	}

	return nil, err
}

// CheckUserStatus TODO: 垃圾实现，需要重构，暂时先这样，token永远自动续（我再研究研究）
func (j *JWTService) CheckUserStatus(user *po.User) (string, error) {
	// 此时用户已经存在，但是没有token，需要生成token
	var token string
	token, err := j.GetKeyFromRedis(user.Uuid)
	// 未能获取到就生成token
	if err != nil {
		token, err = j.GenerateJWTToken(user.UserName, user.Password)
	}
	_, err = j.ParseToken(token)
	if err != nil {
		log.Panicln(constant.LoginErrorTokenExpired.Message)
	}
	return token, err
}
