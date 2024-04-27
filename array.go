package main

import "fmt"

func main() {
	var student = [5]string{"rahim", "karim", "jamil", "hamid", "rony"}

	fmt.Println("student list", student)

	fmt.Println("student1", student[0])

	fmt.Println("total student", len(student))
}
