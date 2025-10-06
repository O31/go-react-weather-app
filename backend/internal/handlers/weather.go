package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"weather-backend/internal/services"

	"github.com/go-chi/chi/v5"
)

func GetWeatherByCity(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")

	if city == "" {
		http.Error(w, "City parameter is required", http.StatusBadRequest)
		return
	}

	weather := services.GetWeatherByCity(strings.Title(strings.ToLower(city)))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
