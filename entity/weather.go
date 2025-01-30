package entity

import "errors"

type Weather struct {
	Temperature float64 `json:"temperature"`
	WindSpeed   float64 `json:"windspeed"`
	WeatherCode int     `json:"weathercode"`
}

var ErrWeatherNotFound = errors.New("weather data not found")

func NewWeather(temperature, windSpeed float64, weatherCode int) Weather {
	return Weather{
		Temperature: temperature,
		WindSpeed:   windSpeed,
		WeatherCode: weatherCode,
	}
}

func (w *Weather) IsRainy() bool {
	return w.WeatherCode >= 200 && w.WeatherCode < 600
}

func (w *Weather) IsSunny() bool {
	return w.WeatherCode == 800
}

func (w *Weather) IsCloudy() bool {
	return w.WeatherCode > 800 && w.WeatherCode < 900
}
