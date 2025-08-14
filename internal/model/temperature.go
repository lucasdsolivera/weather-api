package model

type Temperature struct {
	Kelvin     float64 `json:"kelvin"`
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
}
