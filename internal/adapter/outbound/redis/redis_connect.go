package redis

import (
	"Go-Blog/config"
	"fmt"
	"github.com/go-redis/redis"
)

func CreateRedisConnect() *redis.Client {
	host := config.GetYamlConfig().Redis.Host
	port := config.GetYamlConfig().Redis.Port
	fmt.Println(host + ":" + port)
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port, // redis地址
		Password: "",                // redis密码，没有则留空
		DB:       0,                 // 默认数据库，默认是0
	})
	return client
}
