package main

import (
	"andreani/goweather/api"
	"fmt"
	"sync"
)

func main() { // main routine
	dummyCities := []string{"Madrid", "Buenos Aires", "Pilar"}
	var wg sync.WaitGroup
	for _, city := range dummyCities {
		wg.Add(1)
		go func(city string) {
			renderWeather(city)
			wg.Done()
		}(city)
	}
	wg.Wait() // wait synchronously to the whole group
	// time.Sleep(3 * time.Second)  Hack
}

func renderWeather(city string) {
	fmt.Println("Fetching weather for ", city, " ⌛️")
	wc, err := api.GetWeather(city)
	if err != nil {
		fmt.Println("We couldn't fetch the weather: ", err.Error())
	} else {
		fmt.Printf("%v %.1f⁰C \n", wc.Name, wc.Temperature.ToCelsius())
	}
}
