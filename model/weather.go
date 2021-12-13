package model

type Temperature float64

// func (t Temperature) toFahrenheit() float64 {
// 	return float64(t) * 2
// }

type Weather struct {
	temperature Temperature
	feelsLike   Temperature
	condition   int
	visibility  float64
}
