package main

import "fmt"

func main() {
	var userInput int

	fmt.Println("enter your number")
	fmt.Scan(&userInput)

	if userInput == 0 {
		fmt.Printf("%v User input is zero", userInput)
	} else if userInput%2 == 0 {
		fmt.Printf("%v User input is even", userInput)
	} else {
		fmt.Printf("%v User input is odd", userInput)
	}
}
