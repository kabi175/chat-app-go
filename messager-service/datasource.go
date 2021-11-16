package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresClient() (*gorm.DB, error) {
	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=5432 sslmode=disable TimeZone=Asia/Shanghai", POSTGRES_HOST, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func NewRedisClient() *redis.Client {

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
