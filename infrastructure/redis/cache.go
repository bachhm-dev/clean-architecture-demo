package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bachhm.dev/clean-architecture-service/service"
	"github.com/bachhm.dev/clean-architecture-service/service/entity"

	"github.com/go-redis/redis/v8"
)

type weatherRepository struct {
	Client *redis.Client
}

func NewRedisRepository(client *redis.Client) service.WeatherRepository {
	return weatherRepository{Client: client}
}

func (r weatherRepository) GetWeather(ctx context.Context, latitude, longitude float64) (*entity.Weather, error) {
	key := fmt.Sprintf("weather:%f:%f", latitude, longitude)
	data, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("weather data not found in cache")
	}
	if err != nil {
		return nil, err
	}

	var weather entity.Weather
	err = json.Unmarshal([]byte(data), &weather)
	if err != nil {
		return nil, err
	}

	return &weather, nil
}

func (r weatherRepository) SaveWeather(ctx context.Context, latitude, longitude float64, weather *entity.Weather) error {
	key := fmt.Sprintf("weather:%f:%f", latitude, longitude)
	data, err := json.Marshal(weather)
	if err != nil {
		return err
	}

	// Cache the data for 1 hour
	err = r.Client.Set(ctx, key, data, time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}
