package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storageService = &StorageService{}
	ctx            = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("couldn't connect to redis %v", err))
	}
	fmt.Printf("connection to redis successful: %v", pong)
	storageService.redisClient = redisClient
	return storageService
}

func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storageService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

func RetrieveInitialUrl(shortUrl string) string {
	url, err := storageService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return url
}
