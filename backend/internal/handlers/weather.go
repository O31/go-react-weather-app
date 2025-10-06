package handlers

import (
	"encoding/json"
	"net/http"

	"weather-backend/internal/services"

	"github.com/go-chi/chi"
)

func GetWeatherByCity(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city") // Get the {city} from the route

	if city == "" {
		city = services.GetLastLocation(r)
	}

	weather := services.GetWeatherByCity(city)

	services.SetLastLocation(w, weather.City)
	services.AddRecentSearch(w, r, weather.City)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}

func RecentSearchesHandler(w http.ResponseWriter, r *http.Request) {
	recentSearches := services.GetRecentSearches(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recentSearches)
}
