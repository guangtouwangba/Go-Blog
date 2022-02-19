package service

import (
	"Go-Blog/internal/constant"
	uuid "github.com/satori/go.uuid"
	"time"
)

type JWTService struct {
}

func GetKeyFromRedis(userId uuid.UUID) (string, error) {
	jwt, err := constant.RedisConnect.Get(userId.String()).Result()
	return jwt, err
}

func SetKeyToRedis(userId uuid.UUID, jwt string, duration time.Duration) (string, error) {
	value, err := constant.RedisConnect.Set(userId.String(), jwt, duration*time.Second).Result()
	return value, err
}
