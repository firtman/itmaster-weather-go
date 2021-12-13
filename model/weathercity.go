package model

type WeatherCity struct {
	// composition
	// Weather Weather
	// City    City

	// embedding
	Weather
	City
	timestamp string
}
