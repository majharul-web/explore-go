package main

import "fmt"

func main() {

	var numbers = []int{1, 2, 3, 4, 5, 6}

	for i := range numbers {
		fmt.Println("Slice item", i, "is", numbers[i])
	}

	/* create a map*/
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}

	/* print map using keys*/
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* print map using key-value*/
	for country, capital := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", capital)
	}
}
