package main

import (
	"andreani/goweather/model"
	"fmt"
)

func main() {
	caba := model.WeatherCity{}
	caba.Name = "CABA"
	caba.Country = "AR"
	fmt.Printf("La ciudad es: %v", caba)
}
