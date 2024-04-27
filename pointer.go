package main

import "fmt"

func main() {

	var x int = 10

	var p *int
	p = &x

	fmt.Println("value of x is = ", x)
	fmt.Println("address of x is = ", &x)
	fmt.Println("value of pointer is = ", p)
	fmt.Println("value of pointer is x = ", *p)

	*p = *p - 1
	fmt.Println("after decrease", x)
}
