package main

import (
	"fmt"
)

type student struct {
	name       string
	age        int
	university string
}

func main() {
	student1 := student{"John", 18, "UCL"}
	fmt.Println(student1)

	student2 := student{
		name:       "Mark",
		age:        19,
		university: "Brighton University",
	}

	fmt.Println(student2)

	student3 := student{name: "Clara", age: 20}
	fmt.Println(student3)

	student3.university = "Birmingham University"
	fmt.Println(student3)
}
