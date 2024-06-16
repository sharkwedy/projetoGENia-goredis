package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

type objectService struct {
	redisClient *redis.Client
	ctx         context.Context
}

func NewObjectService(redisClient *redis.Client, ctx context.Context) ObjectService {
	return &objectService{
		redisClient: redisClient,
		ctx:         ctx,
	}
}

func (s *objectService) FetchObjects() ([]Object, error) {
	const cacheKey = "objects"

	// Try to get data from Redis cache
	cachedObjects, err := s.redisClient.Get(s.ctx, cacheKey).Result()
	if err == nil {
		var objects []Object
		if jsonErr := json.Unmarshal([]byte(cachedObjects), &objects); jsonErr == nil {
			return objects, nil
		}
		log.Printf("Failed to unmarshal cached data: %v", jsonErr)
	}

	// If not found in cache, fetch from external API
	resp, err := http.Get("https://api.restful-api.dev/objects")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var objects []Object
	if err := json.Unmarshal(body, &objects); err != nil {
		return nil, err
	}

	// Store the fetched objects in Redis cache
	if jsonData, err := json.Marshal(objects); err == nil {
		s.redisClient.Set(s.ctx, cacheKey, jsonData, 0)
	}

	return objects, nil
}
