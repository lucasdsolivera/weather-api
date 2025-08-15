package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type WeatherAPI interface {
	FetchLocation(city, state, country string) ([]byte, error)
	FetchTemperature(lat, lon float64) ([]byte, error)
}

type OpenWeatherAPIClient struct {
	BaseURL string
	APIKey  string
	Client  *http.Client
}

func NewAPIClient() *OpenWeatherAPIClient {
	baseURL := os.Getenv("OPEN_WEATHER_API_URL")
	if baseURL == "" {
		panic("OPEN_WEATHER_API_URL not set")
	}

	apiKey := os.Getenv("OPEN_WEATHER_API_KEY")
	if apiKey == "" {
		panic("OPEN_WEATHER_API_KEY not set")
	}

	return &OpenWeatherAPIClient{
		BaseURL: baseURL,
		APIKey:  apiKey,
		Client:  &http.Client{Timeout: 3 * time.Second},
	}
}

func (c *OpenWeatherAPIClient) FetchLocation(city, state, country string) ([]byte, error) {
	cityEsc := url.QueryEscape(city)
	stateEsc := url.QueryEscape(state)
	countryEsc := url.QueryEscape(country)

	url := fmt.Sprintf(
		"%s/geo/1.0/direct?q=%s,%s,%s&limit=10&appid=%s", c.BaseURL, cityEsc, stateEsc, countryEsc, c.APIKey,
	)

	logURL := strings.Replace(url, c.APIKey, "***", 1)
	log.Printf("Fetching location: %s", logURL)

	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, NewHTTPError(resp.StatusCode, string(body))
	}

	return io.ReadAll(resp.Body)
}

func (c *OpenWeatherAPIClient) FetchTemperature(lat, lon float64) ([]byte, error) {
	url := fmt.Sprintf(
		"%s/data/2.5/weather?lat=%f&lon=%f&appid=%s",
		c.BaseURL, lat, lon, c.APIKey,
	)

	logURL := strings.Replace(url, c.APIKey, "***", 1)
	log.Printf("Fetching temperature: %s", logURL)

	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, NewHTTPError(resp.StatusCode, string(body))
	}

	return io.ReadAll(resp.Body)
}
