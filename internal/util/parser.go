package util

import (
	"encoding/json"
	"fmt"

	"github.com/lucasdsolivera/weather-api/internal/model"
)

type mainData struct {
	Temp float64 `json:"temp"`
}

func ParseTemperature(data []byte, err error) (*model.Temperature, error) {
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

	return &model.Temperature{
		Kelvin:     result.Main.Temp,
		Celsius:    c,
		Fahrenheit: f,
	}, nil
}

func ParseFirstLocation(data []byte, err error) (*model.Location, error) {
	if err != nil {
		return nil, err
	}

	var results []model.Location
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("location not found")
	}

	return &results[0], nil
}
