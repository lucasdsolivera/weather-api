package service

import (
	"encoding/json"

	"github.com/lucasdsolivera/weather-api/internal/client"
	"github.com/lucasdsolivera/weather-api/internal/model"
	"github.com/lucasdsolivera/weather-api/internal/util"
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
	loc, err := util.ParseFirstLocation(s.Client.FetchLocation(city, state, country))
	if err != nil {
		return nil, err
	}

	temp, err := util.ParseTemperature(s.Client.FetchTemperature(loc.Lat, loc.Lon))
	if err != nil {
		return nil, err
	}

	response := struct {
		Location    *model.Location    `json:"location"`
		Temperature *model.Temperature `json:"temperature"`
	}{
		Location:    loc,
		Temperature: temp,
	}

	return json.MarshalIndent(response, "", "  ")
}
