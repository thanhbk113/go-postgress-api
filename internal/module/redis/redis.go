package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"thanhbk113/internal/config"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func InitRedis() {
	ctx := context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.GetConfig().RedisUri,
		Password: config.GetConfig().RedisPassword,
		DB:       0, // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Cannot connect to redis:", config.GetConfig().RedisUri, err)

	}

	fmt.Println(aurora.Green("*** CONNECTED TO REDIS: " + config.GetConfig().RedisUri + " ***"))

}

func GetRedis() *redis.Client {
	return rdb
}

// SetKeyValue ...
func SetKeyValue(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	storeByte, _ := json.Marshal(value)
	r := rdb.Set(ctx, key, storeByte, expiration)
	return r.Err()
}

// GetKeyValue ...
func GetKeyValue(key string, value interface{}) error {
	ctx := context.Background()
	r := rdb.Get(ctx, key)
	if r.Err() != nil {
		return r.Err()
	}
	storeByte, _ := r.Bytes()
	return json.Unmarshal(storeByte, value)
}

// DeleteKey ...
func DeleteKey(key string) error {
	ctx := context.Background()
	r := rdb.Del(ctx, key)
	return r.Err()
}

// GetValues ...
func GetValue(key string) (string, error) {
	ctx := context.Background()
	r := rdb.Get(ctx, key)
	if r.Err() != nil {
		return "", r.Err()
	}
	return r.Val(), nil
}
