package db

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

var ctx = context.Background()

func InitRedisDB(ip, pwd string, db int) {

	redisClient = redis.NewClient(&redis.Options{
		Addr:       ip,
		Password:   pwd,
		DB:         db,
		MaxRetries: 5,
	})

	pong, err := redisClient.Ping(ctx).Result()

	if err != nil {
		fmt.Println("connect redis failed")
	} else {
		fmt.Printf("redis ping result: %s\n", pong)
	}
}

func RedisHashSet(key string, fields interface{}) error {

	if err := redisClient.HSet(ctx, key, fields).Err(); err != nil {
		fmt.Printf("%v \n", err)
		return err
	} else {
		//fmt.Printf("%s success \n", key)
		return nil
	}
}

func RedisHashGetAll(key string) map[string]string {

	result, err := redisClient.HGetAll(ctx, key).Result()

	if err != nil {
		fmt.Printf("%v \n", err)
		return nil
	} else {
		return result
	}
}
