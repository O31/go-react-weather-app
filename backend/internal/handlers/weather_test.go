package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetWeatherByCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/weather/Stockholm", nil)
	w := httptest.NewRecorder()

	GetWeatherByCity(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestRecentSearchesHandlerEmpty(t *testing.T) {
	req := httptest.NewRequest("GET", "/weather/recent", nil)
	w := httptest.NewRecorder()

	RecentSearchesHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
}
