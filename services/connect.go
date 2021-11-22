package services

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

type Red struct {
	Client *redis.Client
}

var RD Red

func Connect() {

	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	Check(err)
	fmt.Println(pong)

	RD = Red{
		Client: client,
	}

}
