// Install Redis client for Go:
// Go has a well-maintained Redis client called go-redis. You can install it using:
// go get github.com/go-redis/redis/v8

package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping(context.Background()).Result()

	if err != nil {
		println(err)
	}

	fmt.Println(pong)

	err = rdb.Set(context.Background(), "key", "value", 0).Err()
	if err != nil {
		println(err)
	}
	val, err := rdb.Get(context.Background(), "key").Result()
	if err != nil {
		println(err)
	}
	fmt.Println(val)

}
