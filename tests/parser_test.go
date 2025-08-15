package tests

import (
	"testing"

	"github.com/lucasdsolivera/weather-api/internal/util"
)

func TestParseFirstLocation_Success(t *testing.T) {
	jsonData := `[{"lat":1.23,"lon":4.56,"name":"City","state":"ST","country":"CT"}]`
	loc, err := util.ParseFirstLocation([]byte(jsonData), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if loc.Lat != 1.23 || loc.Lon != 4.56 || loc.City != "City" {
		t.Errorf("unexpected location: %+v", loc)
	}
}

func TestParseFirstLocation_NotFound(t *testing.T) {
	jsonData := `[]`
	_, err := util.ParseFirstLocation([]byte(jsonData), nil)
	if err != util.ErrLocationNotFound {
		t.Errorf("expected ErrLocationNotFound, got %v", err)
	}
}

func TestParseTemperature_Success(t *testing.T) {
	jsonData := `{"main":{"temp":300.0}}`
	temp, err := util.ParseTemperature([]byte(jsonData), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if temp.Kelvin != 300.0 {
		t.Errorf("expected 300K, got %v", temp.Kelvin)
	}
	if temp.Celsius != 26.85 {
		t.Errorf("expected 26.85°C, got %v", temp.Celsius)
	}
	if temp.Fahrenheit != 80.33 {
		t.Errorf("expected 80.33°F, got %v", temp.Fahrenheit)
	}
}

func TestParseTemperature_InvalidJSON(t *testing.T) {
	jsonData := `invalid`
	_, err := util.ParseTemperature([]byte(jsonData), nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
