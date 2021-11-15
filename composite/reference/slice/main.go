package main

import "fmt"

func main() {
	a := make([]string, 2, 3)
	fmt.Println(a)

	b := []string{"a", "b", "c"}
	fmt.Println(b)

	b = append(b, "d")
	fmt.Println(b)

	for _, val := range b {
		fmt.Println(val)
	}
}
