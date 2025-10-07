package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"weather-backend/internal/models"
)

// Helper function to add a search term to recent searches cookie
func AddRecentSearch(w http.ResponseWriter, r *http.Request, searchTerm string) {
	var recent []string

	// Get existing recent searches from cookie
	if cookie, err := r.Cookie("TTWeather_app"); err == nil {
		recent = strings.Split(cookie.Value, "|")
	}

	// Add new search at the beginning and deduplicate
	recent = append([]string{searchTerm}, recent...)
	seen := make(map[string]bool)
	unique := []string{}

	for _, term := range recent {
		if term != "" && !seen[term] {
			unique = append(unique, term)
			seen[term] = true
		}
		if len(unique) >= 5 { // Keep only last 5 searches
			break
		}
	}

	// Set the updated cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "TTWeather_app",
		Value:    strings.Join(unique, "|"),
		Path:     "/",
		MaxAge:   30 * 24 * 60 * 60, // 30 days
		HttpOnly: false,             // Allow frontend to read if needed
		SameSite: http.SameSiteLaxMode,
	})
}

// Helper function to get recent searches from cookie
func GetRecentSearches(r *http.Request) []string {
	cookie, err := r.Cookie("TTWeather_app")
	if err != nil {
		return []string{}
	}

	searches := strings.Split(cookie.Value, "|")

	result := []string{}
	for _, search := range searches {
		if search != "" {
			result = append(result, search)
		}
	}
	return result
}

func GetWeatherByCity(city string) (models.Weather, error) {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	safeCity := url.QueryEscape(city)

	url := fmt.Sprintf(
		"https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no",
		apiKey, safeCity,
	)

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("WeatherAPI error: status=%d, body=%s\n", resp.StatusCode, string(body))
		return models.Weather{}, fmt.Errorf("failed to fetch weather data for %s", city)
	}
	defer resp.Body.Close()

	var apiResp models.WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		fmt.Println("JSON decode error:", err)
		return models.Weather{}, fmt.Errorf("failed to parse weather data for %s", city)
	}

	return models.Weather{
		City:        apiResp.Location.Name,
		Latitude:    apiResp.Location.Lat,
		Longitude:   apiResp.Location.Lon,
		Temperature: apiResp.Current.TempC,
		FeelsLike:   apiResp.Current.Feelslike,
		Description: apiResp.Current.Condition.Text,
		Humidity:    apiResp.Current.Humidity,
		WindSpeed:   apiResp.Current.WindKph,
		WindDir:     apiResp.Current.WindDir,
		Pressure:    apiResp.Current.PressureMb,
		Visibility:  apiResp.Current.VisKm,
		LocalTime:   apiResp.Location.Localtime,
		Icon:        apiResp.Current.Condition.Icon,
	}, nil
}
