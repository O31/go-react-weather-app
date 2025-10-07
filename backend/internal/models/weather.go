package models

type Weather struct {
	City        string  `json:"city"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Temperature float64 `json:"temperature"`
	FeelsLike   float64 `json:"feels_like"`
	Description string  `json:"description"`
	Humidity    int     `json:"humidity"`
	WindSpeed   float64 `json:"wind_speed"`
	WindDir     string  `json:"wind_direction"`
	Pressure    float64 `json:"pressure"`
	Visibility  float64 `json:"visibility"`
	LocalTime   string  `json:"local_time"`
	Icon        string  `json:"icon"`
	Error       string  `json:"error,omitempty"`
}

type WeatherAPIResponse struct {
	Location struct {
		Name      string  `json:"name"`
		Lat       float64 `json:"lat"`
		Lon       float64 `json:"lon"`
		Localtime string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Feelslike float64 `json:"feelslike_c"`
		Condition struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		} `json:"condition"`
		WindKph    float64 `json:"wind_kph"`
		WindDir    string  `json:"wind_dir"`
		Humidity   int     `json:"humidity"`
		PressureMb float64 `json:"pressure_mb"`
		VisKm      float64 `json:"vis_km"`
	} `json:"current"`
}
