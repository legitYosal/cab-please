package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	utils "github.com/usefss/cab-please/identity/utils"
)

var prefixKey string = "surge-demands:"
var postfixTimeKey string = ":time"

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     utils.GetEnv("REDIS_ADDR"),
	Password: "",
	DB:       0,
})

func IncrementOrSetDemand(key string) {
	reformKey := prefixKey + key
	value, err := rdb.Incr(ctx, reformKey).Result()
	if err != nil {
		panic(err)
	}
	if value <= 1 {
		errT := rdb.Set(ctx, reformKey+postfixTimeKey, time.Now().UTC().Unix(), 0).Err()
		if errT != nil {
			panic(err)
		}
	}
}
