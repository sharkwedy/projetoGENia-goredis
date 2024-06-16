package main

import (
	"net/http"
	"log"
	"os"
	"project/internal/handler"
	"project/internal/service"
	"github.com/go-redis/redis/v8"
	"context"
)

var ctx = context.Background()

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB: 0,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	objectService := service.NewObjectService(redisClient, ctx)
	objectHandler := handler.NewObjectHandler(objectService)

	http.Handle("/objects", objectHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
