package service

import (
	"github.com/lucasdsolivera/weather-api/internal/client"
)

type WeatherService struct {
	Client *client.OpenWeatherAPIClient
}

func NewWeatherService() *WeatherService {
	return &WeatherService{
		Client: client.NewAPIClient(),
	}
}

func (s *WeatherService) GetTemperature(city, state, country string) ([]byte, error) {
	return s.Client.FetchLocation(city, state, country)
	//return json.MarshalIndent(nil, "", "  ")
}
