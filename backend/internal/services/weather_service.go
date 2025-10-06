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
	fmt.Println("\n\n\n\n\n", recent)

	// Set the updated cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "TTWeather_app",
		Value:    strings.Join(unique, "|"),
		Path:     "/",
		MaxAge:   30 * 24 * 60 * 60, // 30 days
		HttpOnly: false,             // Allow frontend to read if needed
		SameSite: http.SameSiteLaxMode,
	})
	fmt.Println("Set cookie:", strings.Join(unique, "|"))
}

// Helper function to get recent searches from cookie
func GetRecentSearches(r *http.Request) []string {
	cookie, err := r.Cookie("TTWeather_app")
	if err != nil {
		return []string{}
	}

	searches := strings.Split(cookie.Value, "|")
	fmt.Println("GetRecentSearches:", searches)
	result := []string{}
	for _, search := range searches {
		if search != "" {
			result = append(result, search)
		}
	}
	return result
}

// Helper function to get last searched location from cookie
func GetLastLocation(r *http.Request) string {
	cookie, err := r.Cookie("lastLocation")
	if err != nil {
		return "Stockholm" // Default city
	}
	return cookie.Value
}

// Helper function to set last location cookie
func SetLastLocation(w http.ResponseWriter, location string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "lastLocation",
		Value:    location,
		Path:     "/",
		MaxAge:   30 * 24 * 60 * 60, // 30 days
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
	})
}

func GetWeatherByCity(city string) models.Weather {
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
		return models.Weather{
			City:        city,
			Temperature: 0.0,
			Description: "Error fetching data",
			Humidity:    0,
			WindSpeed:   0.0,
		}
	}
	defer resp.Body.Close()

	var apiResp models.WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		fmt.Println("JSON decode error:", err)
		return models.Weather{City: city, Temperature: 0.0, Description: "Error parsing data"}
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
	}
}
