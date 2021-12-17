package main

import (
	"andreani/goweather/api"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() { // main routine

	flagInputPtr := flag.String("input", "scan", "select how to enter the cities, valid options: standard, scan")
	flagCitiesPtr := flag.String("cities", "", "a comma-separated list of cities to parse (mandatory if you use input=standard)")
	flagOutputPtr := flag.String("output", "standard", "where to render the weather info, valid options: standard, json")
	flag.Parse()

	var outputFunction func(city string)
	switch *flagOutputPtr {
	case "standard":
		outputFunction = renderWeather
	case "json":
		outputFunction = renderWeatherToJson
	}

	switch *flagInputPtr {
	case "standard":
		var wg sync.WaitGroup
		cities := strings.Split(*flagCitiesPtr, ",")
		for _, city := range cities {
			wg.Add(1)
			go func(city string) {
				outputFunction(city)
				wg.Done()
			}(city)
		}
		wg.Wait()
	case "scan":
		var userCity string
		fmt.Print("Enter the city name: ")
		fmt.Scan(&userCity)
		outputFunction(userCity)
	}

	// wait synchronously to the whole group
	// time.Sleep(3 * time.Second)  Hack
}

func renderWeatherToJson(city string) {
	wc, err := api.GetWeather(city)
	if err != nil {
		fmt.Println("We couldn't fetch the weather: ", err.Error())
	} else {
		filename := fmt.Sprintf("output/%v.json", city)
		json, err := json.Marshal(wc)
		if err != nil {
			fmt.Println("Invalid JSON generation")
		} else {
			if _, err := os.Stat("output"); os.IsNotExist(err) {
				os.Mkdir("output", 0)
			}
			err := os.WriteFile(filename, json, 0)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
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
