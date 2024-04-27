package main

import "fmt"

func add(num1 int, num2 int) {
	result := num1 + num2
	fmt.Printf("%v + %v is =\t%v\n", num1, num2, result)

}

func square(val int) int {
	result := val * val
	return result
}

func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func main() {
	add(10, 20)

	var squareResult = square(5)
	fmt.Println("5 square is=", squareResult)

	//	multiple return
	fmt.Println(myFunction(5, "Hello"))

}
