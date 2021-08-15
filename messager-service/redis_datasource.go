package main

import (
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func redisDataSource() *redis.Client {

	REDIS_HOST := os.Getenv("REDIS_HOST")
	REDIS_PASSWORD := os.Getenv("REDIS_PASSWORD")
	REDIS_DB, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		log.Fatalln(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST,
		Password: REDIS_PASSWORD,
		DB:       REDIS_DB,
	})
	return rdb
}
