package main

import "fmt"

func normalChange(val int) {
	val = 30
}
func referenceChange(val *int) {
	*val = 30
}

func main() {
	var x int = 20
	fmt.Println("value of x is", x)

	normalChange(x)
	fmt.Println("normal change", x)

	referenceChange(&x)
	fmt.Println("reference change", x)

}
