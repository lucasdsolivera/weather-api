package routes

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/lucasdsolivera/weather-api/internal/client"
	"github.com/lucasdsolivera/weather-api/internal/service"
	"github.com/lucasdsolivera/weather-api/internal/util"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/retrieve-temperature", indexHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	state := r.URL.Query().Get("state")
	country := r.URL.Query().Get("country")

	svc := service.NewWeatherService()
	data, err := svc.GetTemperature(city, state, country)
	if err != nil {
		if httpErr, ok := err.(*client.HTTPError); ok {
			w.WriteHeader(httpErr.StatusCode)
			return
		}
		if errors.Is(err, util.ErrLocationNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
