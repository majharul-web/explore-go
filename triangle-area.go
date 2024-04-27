package main

import "fmt"

func main() {
	var base, height, result float32

	fmt.Println("enter your base")
	fmt.Scan(&base)

	fmt.Println("enter your height")
	fmt.Scan(&height)

	result = 0.5 * base * height

	fmt.Printf("%v result is", result)

}
