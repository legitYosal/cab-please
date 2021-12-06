package redis

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	utils "github.com/usefss/cab-please/identity/utils"
)

var postfixTimeKey string = ":time"

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     utils.GetEnv("REDIS_ADDR"),
	Password: "",
	DB:       0,
})

func GetRedisClient() *redis.Client {
	return rdb
}

func IncrementOrSetDemand(key string) {
	value, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	if value <= 1 {
		errT := rdb.Set(ctx, key+postfixTimeKey, time.Now().UTC().Unix(), 0).Err()
		if errT != nil {
			panic(err)
		}
	}
}

func GetDemand(key string) uint64 {
	value, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return 0
	} else if err != nil {
		panic(err)
	}
	demands, errC := strconv.ParseUint(value, 10, 64)
	if errC != nil {
		panic(errC)
	}
	return demands
}
