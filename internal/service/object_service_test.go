package service_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"project/internal/service"
)

func TestFetchObjects(t *testing.T) {
	// Setting up mock Redis
	mockRedis, mock := redismock.NewClientMock()
	ctx := context.Background()

	// Mocked data
	mockData := `[{
		"id": "1",
		"name": "Google Pixel 6 Pro",
		"data": {
			"color": "Cloudy White",
			"capacity": "128 GB"
		}
	}]`

	t.Run("fetch from cache", func(t *testing.T) {
		mock.ExpectGet("objects").SetVal(mockData)

		service := service.NewObjectService(mockRedis, ctx)
		objects, err := service.FetchObjects()

		assert.NoError(t, err)
		assert.Len(t, objects, 1)
		assert.Equal(t, "Google Pixel 6 Pro", objects[0].Name)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("fetch from external API", func(t *testing.T) {
		mock.ExpectGet("objects").RedisNil()
		mock.ExpectSet("objects", mockData, 0).SetVal("OK")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockData))
		}))
		defer ts.Close()

		oldURL := "https://api.restful-api.dev/objects"
		service.ExternalAPIURL = ts.URL
		defer func() { service.ExternalAPIURL = oldURL }()

		service := service.NewObjectService(mockRedis, ctx)
		objects, err := service.FetchObjects()

		assert.NoError(t, err)
		assert.Len(t, objects, 1)
		assert.Equal(t, "Google Pixel 6 Pro", objects[0].Name)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}
