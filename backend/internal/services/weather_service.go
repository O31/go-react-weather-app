package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"weather-backend/internal/models"
)

type openWeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

// func GetWeatherByCity(city string) models.Weather {
// 	apiKey := os.Getenv("OPENWEATHER_API_KEY")

// 	url := fmt.Sprintf(
// 		"https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no",
// 		apiKey, city,
// 	)
// 	fmt.Println("\nRequesting URL:", url) // Debug log to check the request URL

// 	resp, err := http.Get(url)
// 	fmt.Println("Response Status:", resp.Body, resp.Status) // Debug log to check response status
// 	fmt.Println()
// 	if err != nil || resp.StatusCode != 200 {
// 		return models.Weather{
// 			City:        city,
// 			Temperature: 0.0,
// 			Humidity:    0,
// 			WindSpeed:   0.0,
// 		}
// 	}
// 	defer resp.Body.Close()
// 	body, _ := io.ReadAll(resp.Body)
// 	fmt.Printf("OpenWeatherMap error: Status: %d, Body: %s\n", resp.StatusCode, string(body))

// 	var owr openWeatherResponse
// 	if err := json.NewDecoder(resp.Body).Decode(&owr); err != nil {
// 		return models.Weather{
// 			City:        city,
// 			Temperature: 0.0,
// 			Humidity:    0,
// 			WindSpeed:   0.0,
// 		}
// 	}

// 	fmt.Println("OpenWeatherMap Response:", owr) // Debug log to check the decoded response
// 	return models.Weather{
// 		City:        city,
// 		Temperature: owr.Main.Temp,
// 		Humidity:    owr.Main.Humidity,
// 		WindSpeed:   owr.Wind.Speed,
// 	}
// }

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
