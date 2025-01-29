package service

import (
	"context"

	"github.com/bachhm.dev/clean-architecture-service/entity"
)

type service struct {
	repository       WeatherRepository
	openMeteoService OpenMeteoService
}

func NewService(repository WeatherRepository, openMeteoService OpenMeteoService) service {
	return service{repository: repository, openMeteoService: openMeteoService}
}

func (w service) GetWeather(ctx context.Context, latitude, longitude float64) (*entity.Weather, error) {
	// First, check if the weather data exists in the cache/database
	cachedWeather, err := w.repository.GetWeather(ctx, latitude, longitude)
	if err == nil {
		return cachedWeather, nil
	}

	// If no cache, fetch from Open-Meteo API
	weather, err := w.openMeteoService.GetWeather(ctx, latitude, longitude)
	if err != nil {
		return nil, err
	}

	// Save it to cache/database for future requests
	err = w.repository.SaveWeather(ctx, latitude, longitude, weather)
	if err != nil {
		return nil, err
	}

	return weather, nil
}
