package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"weather-backend/internal/services"

	"github.com/go-chi/chi/v5"
)

// func GetWeather(w http.ResponseWriter, r *http.Request) {
// 	weather := services.GetDefaultWeather()
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(weather)
// }

func GetWeatherByCity(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")
	fmt.Println("City parameter:", city) // Debug log

	if city == "" {
		http.Error(w, "City parameter is required", http.StatusBadRequest)
		return
	}

	fmt.Println("Fetching weather for city:", city) // Debug log
	weather := services.GetWeatherByCity(strings.Title(strings.ToLower(city)))
	fmt.Println("Fetched", weather) // Debug log

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
