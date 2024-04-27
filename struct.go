package main

import "fmt"

type Student struct {
	name   string
	id     int
	result float32
}

func (x *Student) increaseId(val int) {
	x.id = x.id + val

}

func displayInfo(student Student) {
	fmt.Println("name is", student.name)
	fmt.Println("age is", student.id)

}

func main() {
	student1 := Student{"rahim", 1, 3.95}

	student2 := Student{"karim", 2, 2.95}
	student2.increaseId(2)

	fmt.Println("student 1 is ", student1)

	displayInfo(student2)

}
