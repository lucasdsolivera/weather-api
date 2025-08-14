package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

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
	url := fmt.Sprintf(
		"%s/geo/1.0/direct?q=%s,%s,%s&limit=10&appid=%s", c.BaseURL, city, state, country, c.APIKey,
	)

	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OpenWeatherMap API returned status %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
