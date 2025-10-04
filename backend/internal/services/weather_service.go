package services

import (
	"weather-backend/internal/models"
)

func GetDefaultWeather() models.Weather {
	return models.Weather{
		City:        "Stockholm",
		Temperature: 12.5,
		Description: "Partly cloudy",
		Humidity:    65,
		WindSpeed:   8.2,
	}
}

func GetWeatherByCity(city string) models.Weather {
	// Mock data for different cities
	weatherData := map[string]models.Weather{
		"Stockholm": {
			City:        "Stockholm",
			Temperature: 12.5,
			Description: "Partly cloudy",
			Humidity:    65,
			WindSpeed:   8.2,
		},
		"Gothenburg": {
			City:        "Gothenburg",
			Temperature: 10.8,
			Description: "Light rain",
			Humidity:    78,
			WindSpeed:   12.1,
		},
		"Malmö": {
			City:        "Malmö",
			Temperature: 14.2,
			Description: "Sunny",
			Humidity:    55,
			WindSpeed:   6.5,
		},
	}

	if weather, exists := weatherData[city]; exists {
		return weather
	}

	// Default fallback
	return models.Weather{
		City:        city,
		Temperature: 15.0,
		Description: "Clear sky",
		Humidity:    60,
		WindSpeed:   7.0,
	}
}
