package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"project/internal/service"
	"project/internal/handler"
)

func TestObjectHandler(t *testing.T) {
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

	mock.ExpectGet("objects").SetVal(mockData)

	// Initialize the service and handler
	objectService := service.NewObjectService(mockRedis, ctx)
	objectHandler := handler.NewObjectHandler(objectService)

	r := httptest.NewRequest(http.MethodGet, "/objects", nil)
	w := httptest.NewRecorder()

	// Serve the request
	objectHandler.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	// Checking the response
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Check the response body
	var responseObject []service.Object
	err := json.NewDecoder(res.Body).Decode(&responseObject)
	assert.NoError(t, err)
	assert.Len(t, responseObject, 1)
	assert.Equal(t, "Google Pixel 6 Pro", responseObject[0].Name)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}
