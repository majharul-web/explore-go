package main

import "fmt"

func main() {
	var car = map[string]string{"model": "bmw", "quantity": "5"}
	country := map[string]string{"country1": "bangladesh", "country2": "japan"}

	delete(country, "country1")

	fmt.Println("car", car)
	fmt.Println("country", country)
}
