package service

import (
	"context"

	"github.com/bachhm.dev/clean-architecture-service/service/entity"
)

type WeatherUsecase interface {
	GetWeather(ctx context.Context, latitude, longitude float64) (*entity.Weather, error)
}

type WeatherRepository interface {
	GetWeather(ctx context.Context, latitude, longitude float64) (*entity.Weather, error)
	SaveWeather(ctx context.Context, latitude, longitude float64, weather *entity.Weather) error
}

type OpenMeteoService interface {
	GetWeather(ctx context.Context, latitude, longitude float64) (*entity.Weather, error)
}
