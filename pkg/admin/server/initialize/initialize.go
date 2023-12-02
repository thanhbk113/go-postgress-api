package initialize

import (
	"thanhbk113/internal/config"
	"thanhbk113/internal/module/redis"
	"thanhbk113/pkg/admin/server/initialize/kafka"
)

func Init() {
	config.Init()
	database()
	redis.InitRedis()
	kafka.InitKafka()

}
