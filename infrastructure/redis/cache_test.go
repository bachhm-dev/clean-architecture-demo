package redis

import (
	"fmt"
	"testing"

	"context"

	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

// Function that interacts with Redis
func SetValue(ctx context.Context, rdb *redis.Client, key, value string) error {
	return rdb.Set(ctx, key, value, 0).Err()
}

// Unit Test for GetWeather
func TestGetWeather(t *testing.T) {
	db, mock := redismock.NewClientMock()
	ctx := context.Background()

	latitude := 37.7749
	longitude := -122.4194
	key := fmt.Sprintf("weather:%f:%f", latitude, longitude)
	weatherData := `{"temperature": 20.5, "windspeed": 60}`

	// Mock Redis GET command
	mock.ExpectGet(key).SetVal(weatherData)

	repo := NewRedisRepository(db)
	weather, err := repo.GetWeather(ctx, latitude, longitude)
	assert.NoError(t, err)
	assert.NotNil(t, weather)
	assert.Equal(t, 20.5, weather.Temperature)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
