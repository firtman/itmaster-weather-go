package model

import "fmt"

type City struct {
	Id      int
	Name    string
	Country string
}

// método del tipo City, que se llama String y devuelve un string
func (city City) String() string {
	return fmt.Sprintf("[City] %v, %v", city.Name, city.Country)
}
