package cache

import (
	"context"
	"counter/utils"
	"encoding/json"
	"errors"
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

type Value struct {
	Value int
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
	c := Counter{
		Domain: domain,
		Key: key,
		Value: 0,
		AuthKey: utils.GenerateAuthKey(),
	}

	jsonData, err := json.Marshal(c)
	if err != nil {
		return Counter{}, err
	}

	err = Rdb.Set(ctx, utils.CreateRedisKey(domain, key), jsonData, 0).Err()
	if err != nil {
		return Counter{}, err
	}


	return c, nil
}

func Incr(rk string) (Value, error) {
	exists, err := Rdb.Exists(ctx, rk).Result()
	if err != nil {
		return Value{}, err
	}

	if exists == 0 {
		return Value{}, errors.New("Counter not found")
	}

	// Increment the value field by one
	n, err := Rdb.HIncrBy(ctx, rk, "value", 1).Result()
	if err != nil {
		return Value{}, err
	}

	return Value{Value: int(n)}, nil

}

func IncrementDomain(domain, key string) Value {
	redisKey := fmt.Sprintf("counter:%s:%s", domain, key)

	var val int


	
	// increment with domain and key (unique) and domain is optional

	data := Rdb.Get(ctx, redisKey)

	fmt.Println(data)

	return Value{
		Value: val + 1,
	}
	
}
