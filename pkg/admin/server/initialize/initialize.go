package initialize

import (
	"thanhbk113/internal/config"
	"thanhbk113/internal/module/redis"
)

func Init() {
	config.Init()
	database()
	redis.InitRedis()
	initKafka()

}
