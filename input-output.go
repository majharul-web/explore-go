package main

import "fmt"

func main() {

	//	input--->Scan,Scanf,Scanln
	//	output--->Print,Printf,Println

	var num1, num2 int

	fmt.Println("Enter your 2 number")
	fmt.Scanf("%v %v", &num1, &num2)

	fmt.Printf("Num1 is=%v and num2 is=%v", num1, num2)
}
