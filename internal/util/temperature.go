package util

import "math"

func round(value float64, places int) float64 {
	factor := math.Pow(10, float64(places))
	return math.Round(value*factor) / factor
}

func ConvertTemperature(k float64) (c, f float64) {
	c = round(k-273.15, 2)
	f = round((k-273.15)*9/5+32, 2)
	return
}
