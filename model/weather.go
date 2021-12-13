package model

type Temperature float64

func (t Temperature) ToCelsius() float64 {
	return float64(t) - 273.15
}

type Weather struct {
	Temperature Temperature
	FeelsLike   Temperature
	Condition   int
	Visibility  float64
}
