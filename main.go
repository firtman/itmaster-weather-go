package main

import (
	"andreani/goweather/api"
	"fmt"
)

func main() {
	wc, err := api.GetWeather("Madrid")
	if err != nil {
		fmt.Println("We couldn't fetch the weather: ", err.Error())
	} else {
		fmt.Println(wc)
	}
}
