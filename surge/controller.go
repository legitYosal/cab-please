package main

import (
	bgContext "context"
	"encoding/json"

	redis "github.com/go-redis/redis/v8"
	redisClient "github.com/usefss/cab-please/passenger/redis"
	"github.com/usefss/cab-please/surge/stubs"
	"golang.org/x/net/context"
)

type Threshold struct {
	RequestThreshold uint64  `json:"request_threshold"`
	PriceCoefficient float32 `json:"price_coefficient"`
}

var SURGE_CONFIG_CACHE_KEY string = "surge-service-config"

type GRPCServer struct {
}

func getConfigs() *[]Threshold {
	ctx := bgContext.Background()
	rawConfig, err := redisClient.GetRedisClient().Get(ctx, SURGE_CONFIG_CACHE_KEY).Result()
	if err == redis.Nil {
		panic("No surge config is in cache")
	} else if err != nil {
		panic(err)
	}
	var configs []Threshold
	errJson := json.Unmarshal([]byte(rawConfig), &configs)
	if errJson != nil {
		panic(errJson)
	}
	return &configs
}

func (s *GRPCServer) GetSurgeRate(ctx context.Context, in *stubs.DemandAreaKey) (*stubs.SurgeRate, error) {
	demandCount := redisClient.GetDemand(in.Key)
	configs := *getConfigs()
	for _, threshold := range configs {
		if demandCount >= threshold.RequestThreshold {
			return &stubs.SurgeRate{Rate: threshold.PriceCoefficient}, nil
		}
	}
	panic("No surge rate found")
}
