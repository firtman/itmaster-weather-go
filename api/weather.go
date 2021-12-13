package api

import (
	"andreani/goweather/model"
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
	//TODO: parsear JSON
	wc.Name = "Dummy"
	wc.Country = "AL"
	wc.Temperature = 31
}
