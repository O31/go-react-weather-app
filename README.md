# Weather App

A modern, interactive weather dashboard combining a **React** frontend and a **Go** backend. Search
for weather in any city, click directly on the map to view local conditions, and enjoy personalized
session features with cookies. Proeject was done to learn about accessing API with Go and to
integrate API requests from go to a React frontend.

## Features

- **Current Weather:** Live data for any city worldwide.
- **Interactive Map:** Click anywhere to view weather by location.
- **Recent Searches:** Last searched cities always at your fingertips.
- **Session Persistence:** Keeps your last location and history across visits.
- **Comprehensive Weather Details:** Temperature, feels like, humidity, wind, pressure, visibility,
  and local time.
- **Responsive Design:** Seamless experience on desktop and mobile.
- **Loading & Error Animations:** Clean, user-friendly interface.
- **Cookie-Based Storage:** Session features without requiring an account.

## Getting Started

### Prerequisites

- [Node.js](https://nodejs.org/) & npm
- [Go](https://golang.org/) (v1.19+)
- [WeatherAPI](https://www.weatherapi.com/) API key

### Backend Setup (Go)

create a .env file and add your weatherapi key as OPENWEATHER_API_KEY=key <br>cd backend <br> go run
cmd/server/main.go <br>

### Frontend Setup (React)

cd frontend <br> npm install <br> npm run dev <br>

## Usage

- **Search**: Enter a city name to view weather details.
- **Map**: Click anywhere on the interactive map for live weather.
- **Recent Searches**: Click any previous city for quick access.
- **Session Persistence**: Return to the app and continue from your last location.

## Customization Ideas

- 5-day or hourly forecast
- "My location" button using geolocation
- Add favorite/pinned cities
- Dark/light mode toggle
- Autocomplete search
