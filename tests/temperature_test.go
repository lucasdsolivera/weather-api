package tests

import (
	"testing"

	"github.com/lucasdsolivera/weather-api/internal/util"
)

func TestConvertTemperature(t *testing.T) {
	tests := []struct {
		kelvin    float64
		expectedC float64
		expectedF float64
	}{
		{0, -273.15, -459.67},
		{273.15, 0, 32},
		{300, 26.85, 80.33},
	}

	for _, tt := range tests {
		c, f := util.ConvertTemperature(tt.kelvin)
		if c != tt.expectedC {
			t.Errorf("Kelvin %.2f: expected Celsius %.2f, got %.2f", tt.kelvin, tt.expectedC, c)
		}
		if f != tt.expectedF {
			t.Errorf("Kelvin %.2f: expected Fahrenheit %.2f, got %.2f", tt.kelvin, tt.expectedF, f)
		}
	}
}
