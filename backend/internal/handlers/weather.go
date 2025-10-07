package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"weather-backend/internal/models"
	"weather-backend/internal/services"

	"github.com/go-chi/chi/v5"
)

func GetWeatherByCity(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")

	if city == "" {
		fmt.Println("City param empty, using recent search if available")
		if len(services.GetRecentSearches(r)) != 0 {
			city = services.GetRecentSearches(r)[0]
		} else {
			city = "Stockholm" // Default city
		}
	}
	fmt.Println("GetWeatherByCity:", city)

	weather, err := services.GetWeatherByCity(city)
	if err != nil {
		json.NewEncoder(w).Encode(models.Weather{
			Error: "Invalid city name or API error",
		})
		return
	}

	services.AddRecentSearch(w, r, weather.City)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}

func RecentSearchesHandler(w http.ResponseWriter, r *http.Request) {
	recentSearches := services.GetRecentSearches(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recentSearches)
}
