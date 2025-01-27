package open_meteo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bachhm.dev/clean-architecture-service/service"
	"github.com/bachhm.dev/clean-architecture-service/service/entity"
)

type openMeteoAPI struct{}

func New() service.OpenMeteoService {
	return openMeteoAPI{}
}

const openMeteoEndpoint = "https://api.open-meteo.com/v1/forecast"

func (api openMeteoAPI) GetWeather(ctx context.Context, latitude, longitude float64) (*entity.Weather, error) {
	url := fmt.Sprintf("%s?latitude=%f&longitude=%f&current_weather=true", openMeteoEndpoint, latitude, longitude)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch weather data: %v", resp.Status)
	}

	var result struct {
		CurrentWeather entity.Weather `json:"current_weather"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result.CurrentWeather, nil
}
