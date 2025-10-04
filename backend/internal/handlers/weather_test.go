package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestGetWeather(t *testing.T) {
	req, err := http.NewRequest("GET", "/weather", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetWeather)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if response contains expected fields
	body := rr.Body.String()
	if !strings.Contains(body, "Stockholm") {
		t.Errorf("Expected Stockholm in response, got %v", body)
	}
	if !strings.Contains(body, "temperature") {
		t.Errorf("Expected temperature field in response, got %v", body)
	}
}

func TestGetWeatherByCity(t *testing.T) {
	// Create a Chi router for proper URL parameter handling
	r := chi.NewRouter()
	r.Get("/weather/{city}", GetWeatherByCity)

	req, err := http.NewRequest("GET", "/weather/Gothenburg", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "Gothenburg") {
		t.Errorf("Expected Gothenburg in response, got %v", body)
	}
}
