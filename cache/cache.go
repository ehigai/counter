package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client
var ctx context.Context

type Counter struct {
    Domain  string `json:"domain"`
    Key     string `json:"key"`
    Value   int    `json:"value"`
    AuthKey string `json:"authKey"`
}

func redisConnectionOptions() *redis.Options {
	options := &redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0, // Use Default database
		Protocol: 2,
	}
	return options
}

func init() {
	Rdb = redis.NewClient(redisConnectionOptions())
	ctx = context.Background()

}

func CreateCounter(domain, key string) (Counter, error) {
	newCounter := Counter{
		Domain: domain,
		Key: key,
		Value: 0,
		AuthKey: "auto-generated",
	}

	jsonData, err := json.Marshal(newCounter)
	if err != nil {
		return Counter{}, err
	}

	redisKey := fmt.Sprintf("counter:%s:%s", domain, key)

	err = Rdb.Set(ctx, redisKey, jsonData, 0).Err()
	if err != nil {
		return Counter{}, err
	}


	return newCounter, nil
}
