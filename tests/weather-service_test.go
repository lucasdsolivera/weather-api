package tests

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/lucasdsolivera/weather-api/internal/model"
	"github.com/lucasdsolivera/weather-api/internal/service"
	"github.com/lucasdsolivera/weather-api/internal/util"
)

type mockClient struct {
	fetchLocationData  []byte
	fetchLocationError error
	fetchTempData      []byte
	fetchTempError     error
}

func (m *mockClient) FetchLocation(city, state, country string) ([]byte, error) {
	return m.fetchLocationData, m.fetchLocationError
}

func (m *mockClient) FetchTemperature(lat, lon float64) ([]byte, error) {
	return m.fetchTempData, m.fetchTempError
}

func TestGetTemperature_Success(t *testing.T) {
	mock := &mockClient{
		fetchLocationData: []byte(`[{"lat":1.23,"lon":4.56,"name":"TestCity","state":"TS","country":"TC"}]`),
		fetchTempData:     []byte(`{"main":{"temp":300}}`),
	}

	svc := &service.WeatherService{
		Client: mock,
	}

	data, err := svc.GetTemperature("TestCity", "TS", "TC")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var resp struct {
		Location    *model.Location    `json:"location"`
		Temperature *model.Temperature `json:"temperature"`
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if resp.Location.City != "TestCity" || resp.Temperature.Kelvin != 300 {
		t.Errorf("unexpected response: %+v", resp)
	}
}

func TestGetTemperature_LocationNotFound(t *testing.T) {
	mock := &mockClient{
		fetchLocationData: []byte(`[]`),
	}

	svc := &service.WeatherService{
		Client: mock,
	}

	_, err := svc.GetTemperature("Nowhere", "XX", "YY")
	if !errors.Is(err, util.ErrLocationNotFound) {
		t.Errorf("expected ErrLocationNotFound, got %v", err)
	}
}

func TestGetTemperature_FetchTemperatureError(t *testing.T) {
	locJSON := `[{"lat":1.23,"lon":4.56,"name":"TestCity","state":"TS","country":"TC"}]`

	mock := &mockClient{
		fetchLocationData: []byte(locJSON),
		fetchTempError:    errors.New("network error"),
	}

	svc := &service.WeatherService{
		Client: mock,
	}

	_, err := svc.GetTemperature("TestCity", "TS", "TC")
	if err == nil || err.Error() != "network error" {
		t.Errorf("expected network error, got %v", err)
	}
}
