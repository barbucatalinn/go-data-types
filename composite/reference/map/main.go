package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1)

	m2 := make(map[string]int)
	fmt.Println(m2)

	m3 := map[string]int{
		"a": 18,
		"b": 25,
	}
	fmt.Println(m3)

	// add
	m3["c"] = 35

	// get
	fmt.Printf("%d\n", m3["a"])

	// delete
	delete(m3, "b")

	fmt.Println(m3)
}
