import React, { useState, useEffect } from "react"
import WeatherMap from "./WeatherMap"
import "./Weather.css"

const WEATHER_API_URL = "http://localhost:8080/weather"

function Weather() {
  const [city, setCity] = useState("Stockholm")
  const [weather, setWeather] = useState(null)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState("")
  const [recentSearches, setRecentSearches] = useState([])
  const [selectedLocation, setSelectedLocation] = useState({ lat: 59.3293, lng: 18.0686 }) // Stockholm default

  // Load recent searches from localStorage on component mount
  useEffect(() => {
    const saved = localStorage.getItem("TTWeather_app")
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

  const fetchWeather = async (query) => {
    setError("")
    setLoading(true)
    try {
      const res = await fetch(`${WEATHER_API_URL}/${query}`, {
        credentials: "include",
      })
      console.log("Fetch response: ", `${WEATHER_API_URL}/${query}`)
      if (!res.ok) throw new Error("City not found or API error")
      const data = await res.json()
      if (data.error !== undefined) {
        setWeather(null)
        throw new Error("Error fetching data from API")
      }

      fetchRecentSearches()
      setWeather(data)
      setSelectedLocation({ lat: data.latitude, lng: data.longitude })
      setCity(data.city)
    } catch (err) {
      setError(err.message)
      setWeather(null)
    } finally {
      setLoading(false)
    }
  }

  // Handle map location selection
  const handleLocationSelect = async (lat, lng) => {
    const locationQuery = `${lat.toFixed(4)},${lng.toFixed(4)}`
    fetchWeather(locationQuery)
  }

  // Add function to fetch recent searches from backend
  const fetchRecentSearches = async () => {
    try {
      const res = await fetch(`${WEATHER_API_URL}/recent`, {
        credentials: "include",
      })
      console.log("res res res:", res)
      if (res.ok) {
        const searches = await res.json()
        console.log("Recent searches fetched:", searches)
        setRecentSearches(searches || [])
      }
    } catch (err) {
      console.log("Could not fetch recent searches:", err)
    }
  }

  // Update useEffect to fetch recent searches on mount
  useEffect(() => {
    fetchRecentSearches()
    fetchWeather("")
  }, [])

  const handleSearch = async (e) => {
    e.preventDefault()
    if (!city.trim()) return
    console.log("Searching for city:", city)
    fetchWeather(city.trim())
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
              <span key={index} className="recent-city" onClick={() => fetchWeather(recentCity)}>
                {recentCity}
              </span>
            ))}
          </div>
        )}
      </div>
      {error && <div className="error">{error}</div>}

      {weather && (
        <>
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
        </>
      )}
      <WeatherMap onLocationSelect={handleLocationSelect} selectedLocation={selectedLocation} />
    </div>
  )
}

export default Weather
