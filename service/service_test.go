package service

import (
	"context"
	"testing"

	"github.com/bachhm.dev/clean-architecture-service/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockWeatherRepository is a mock implementation of the WeatherRepository interface
type MockWeatherRepository struct {
	mock.Mock
}

// MockOpenMeteoService is a mock implementation of the OpenMeteoService interface
type MockOpenMeteoService struct {
	mock.Mock
}

func (m *MockWeatherRepository) GetWeather(ctx context.Context, latitude, longitude float64) (*entity.Weather, error) {
	args := m.Called(ctx, latitude, longitude)
	return args.Get(0).(*entity.Weather), args.Error(1)
}

func (m *MockWeatherRepository) SaveWeather(ctx context.Context, latitude, longitude float64, weather *entity.Weather) error {
	args := m.Called(ctx, latitude, longitude, weather)
	return args.Error(0)
}

func (m *MockOpenMeteoService) GetWeather(ctx context.Context, latitude, longitude float64) (*entity.Weather, error) {
	args := m.Called(ctx, latitude, longitude)
	return args.Get(0).(*entity.Weather), args.Error(1)
}

func TestWeatherServiceGetWeatherCacheHit(t *testing.T) {
	mockRepo := new(MockWeatherRepository)
	mockMeteoService := new(MockOpenMeteoService)
	service := NewService(mockRepo, mockMeteoService)

	expectedWeather := &entity.Weather{Temperature: 25.0}
	mockRepo.On("GetWeather", mock.Anything, 10.0, 20.0).Return(expectedWeather, nil)

	ctx := context.Background()
	weather, err := service.GetWeather(ctx, 10.0, 20.0)

	assert.NoError(t, err)
	assert.Equal(t, expectedWeather, weather)
	mockRepo.AssertExpectations(t)
	mockMeteoService.AssertNotCalled(t, "GetWeather", mock.Anything, 10.0, 20.0)
}
