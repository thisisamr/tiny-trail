package db

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Redis_Client *redis.Client

func CreatredisClient(dbn int) {

	Redis_Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: os.Getenv("DB_PASSWORD"), // no password set
		DB:       dbn,                      // use default DB
	})

}
