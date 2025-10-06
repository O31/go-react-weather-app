import React, { useState, useEffect } from "react"
import WeatherMap from "./WeatherMap"
import "./Weather.css"

const WEATHER_API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080/weather"

function Weather() {
  const [city, setCity] = useState("Stockholm")
  const [weather, setWeather] = useState(null)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState("")
  const [recentSearches, setRecentSearches] = useState([])
  const [selectedLocation, setSelectedLocation] = useState({ lat: 59.3293, lng: 18.0686 }) // Stockholm default

  // Load recent searches from localStorage on component mount
  useEffect(() => {
    const saved = localStorage.getItem("recentWeatherSearches")
    if (saved) {
      setRecentSearches(JSON.parse(saved))
    }
  }, [])

  // Auto-clear error after 5 seconds
  useEffect(() => {
    if (error) {
      const timer = setTimeout(() => setError(""), 5000)
      return () => clearTimeout(timer)
    }
  }, [error])

  const addToRecentSearches = (cityName) => {
    const newRecent = [cityName, ...recentSearches.filter((c) => c !== cityName)].slice(0, 5)
    setRecentSearches(newRecent)
    localStorage.setItem("recentWeatherSearches", JSON.stringify(newRecent))
  }

  const fetchWeather = async (cityName) => {
    setLoading(true)
    setError("")
    try {
      const res = await fetch(`${WEATHER_API_URL}/${encodeURIComponent(cityName)}`)
      if (!res.ok) throw new Error("City not found or API error")
      const data = await res.json()
      setWeather(data)
      addToRecentSearches(cityName)
    } catch (err) {
      setError(err.message)
      setWeather(null)
    } finally {
      setLoading(false)
    }
  }

  // Handle map location selection
  const handleLocationSelect = async (lat, lng) => {
    setSelectedLocation({ lat, lng })
    setLoading(true)
    setError("")

    try {
      // Use coordinates for weather query
      const locationQuery = `${lat.toFixed(4)},${lng.toFixed(4)}`
      const res = await fetch(`${WEATHER_API_URL}/${encodeURIComponent(locationQuery)}`)
      if (!res.ok) throw new Error("Weather data not available for this location")
      const data = await res.json()
      setWeather(data)
      setCity(data.city || locationQuery)
      addToRecentSearches(data.city || locationQuery)
    } catch (err) {
      setError(err.message)
      setWeather(null)
    } finally {
      setLoading(false)
    }
  }

  // Initial fetch
  useEffect(() => {
    fetchWeather(city)
    // eslint-disable-next-line
  }, [])

  const handleSearch = (e) => {
    e.preventDefault()
    if (city.trim()) {
      fetchWeather(city.trim())
    }
  }

  const handleRecentCityClick = (recentCity) => {
    setCity(recentCity)
    fetchWeather(recentCity)
  }

  return (
    <div className="weather-card">
      <div className="search-section">
        <form className="search-form" onSubmit={handleSearch}>
          <input
            type="text"
            value={city}
            onChange={(e) => setCity(e.target.value)}
            placeholder="Enter city name"
            disabled={loading}
            autoFocus
          />
          <button type="submit" disabled={loading || !city.trim()}>
            {loading ? "Searching..." : "Search"}
          </button>
        </form>

        {recentSearches.length > 0 && (
          <div className="recent-searches">
            {recentSearches.map((recentCity, index) => (
              <span
                key={index}
                className="recent-city"
                onClick={() => handleRecentCityClick(recentCity)}
              >
                {recentCity}
              </span>
            ))}
          </div>
        )}
      </div>

      {/* Map is always visible now */}
      <WeatherMap onLocationSelect={handleLocationSelect} selectedLocation={selectedLocation} />

      {loading && (
        <div className="loading">
          <div className="spinner"></div>
          Loading weather...
        </div>
      )}

      {error && <div className="error">{error}</div>}

      {weather && !loading && (
        <div className="weather-content">
          <div className="weather-icon">
            <img
              src={weather.icon.startsWith("http") ? weather.icon : `https:${weather.icon}`}
              alt="weather icon"
            />
          </div>
          <div className="weather-info">
            <h3>{weather.city}</h3>
            <div className="temperature">{weather.temperature}°C</div>
            <div className="description">{weather.description}</div>
            <div className="weather-details">
              <div>Feels like: {weather.feels_like}°C</div>
              <div>Humidity: {weather.humidity}%</div>
              <div>
                Wind: {weather.wind_speed} km/h {weather.wind_direction}
              </div>
              <div>Pressure: {weather.pressure} mb</div>
              <div>Visibility: {weather.visibility} km</div>
              <div>Local Time: {weather.local_time}</div>
            </div>
          </div>
        </div>
      )}
    </div>
  )
}

export default Weather
