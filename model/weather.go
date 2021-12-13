package model

type Temperature float64

// func (t Temperature) toFahrenheit() float64 {
// 	return float64(t) * 2
// }

type Weather struct {
	Temperature Temperature
	FeelsLike   Temperature
	Condition   int
	Visibility  float64
}
