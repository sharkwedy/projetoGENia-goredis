package main

import (
	"net/http"
	"log"
	"os"
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

	http.HandleFunc("/objects", func(w http.ResponseWriter, r *http.Request) {
		// Handler code will go here
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
