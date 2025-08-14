package service

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/lucasdsolivera/weather-api/internal/client"
)

type WeatherService struct {
	Client *client.OpenWeatherAPIClient
}

type Location struct {
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	City    string  `json:"name"`
	State   string  `json:"state"`
	Country string  `json:"country"`
}

type Temperature struct {
	Kelvin     float64 `json:"kelvin"`
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
}

type mainData struct {
	Temp float64 `json:"temp"`
}

func NewWeatherService() *WeatherService {
	return &WeatherService{
		Client: client.NewAPIClient(),
	}
}

func (s *WeatherService) GetTemperature(city, state, country string) ([]byte, error) {
	loc, err := parseFirstLocation(s.Client.FetchLocation(city, state, country))
	if err != nil {
		return nil, err
	}

	temp, err := parseTemperature(s.Client.FetchTemperature(loc.Lat, loc.Lon))
	if err != nil {
		return nil, err
	}

	response := struct {
		Location    *Location    `json:"location"`
		Temperature *Temperature `json:"temperature"`
	}{
		Location:    loc,
		Temperature: temp,
	}

	return json.MarshalIndent(response, "", "  ")
}

func parseTemperature(data []byte, err error) (*Temperature, error) {
	if err != nil {
		return nil, err
	}

	var result struct {
		Main mainData `json:"main"`
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	c, f := convertTemperature(result.Main.Temp)

	return &Temperature{
		Kelvin:     result.Main.Temp,
		Celsius:    c,
		Fahrenheit: f,
	}, nil
}

func parseFirstLocation(data []byte, err error) (*Location, error) {
	if err != nil {
		return nil, err
	}

	var results []Location
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("location not found")
	}

	return &results[0], nil
}

func round(value float64, places int) float64 {
	factor := math.Pow(10, float64(places))
	return math.Round(value*factor) / factor
}

func convertTemperature(k float64) (c, f float64) {
	c = round(k-273.15, 2)
	f = round((k-273.15)*9/5+32, 2)
	return
}
