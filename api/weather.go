package api

import (
	"andreani/goweather/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	apiUrl = "https://api.openweathermap.org/data/2.5/weather?q=%v&appid=%v"
	apiKey = "85ad76cb241ddd400e9c40d1b59c5f74"
)

func GetWeather(cityName string) (*model.WeatherCity, error) {
	url := fmt.Sprintf(apiUrl, cityName, apiKey)
	res, err := http.Get(url) // sincrónico cuando los headers están listos
	if err != nil {
		// hubo error de red (el servidor no response, el dns no anda, no hay conexión)
		return nil, err
	}
	if res.StatusCode == http.StatusOK {
		// Estamos OK
		body, err := io.ReadAll(res.Body) // pide leer todo el cuerpo de la respuesta HTTP
		if err != nil {
			// hubo error de red al descargar el body
			return nil, err
		}
		// bodyStr := string(body)
		// fmt.Println(bodyStr)
		wc := model.WeatherCity{}
		parseWeatherJson(body, &wc)
		return &wc, nil
	} else {
		// hubo error en el servidor
		// return nil, errors.New(fmt.Sprintf("Status code error %v", res.Status))
		return nil, fmt.Errorf("status code error %v", res.Status)
	}
}

func parseWeatherJson(bytes []byte, wc *model.WeatherCity) {
	var result OpenWeatherMapResponse
	json.Unmarshal(bytes, &result)

	wc.Name = result.Name
	wc.Id = int(result.Sys.Id)
	wc.Country = result.Sys.Country
	wc.Temperature = model.Temperature(result.Main.Temperature)
}

func parseManualWeatherJson(bytes []byte, wc *model.WeatherCity) {
	var result map[string]interface{}
	json.Unmarshal(bytes, &result)

	wc.Name = result["name"].(string)
	sys := result["sys"].(map[string]interface{})
	wc.Id = int(sys["id"].(float64))
	wc.Country = sys["country"].(string)
	main := result["main"].(map[string]interface{})
	wc.Temperature = model.Temperature(main["temp"].(float64))
}
