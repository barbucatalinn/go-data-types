package main

import "fmt"

func main() {
	a := "this\nthat"
	fmt.Printf("a is: %s\n", a)

	b := `this\nthat`
	fmt.Printf("b is: %s\n", b)

	// character a and b occupies 1 byte each and £ character occupies 2 bytes
	// the final output will 4 bytes
	c := "ab£"
	fmt.Println([]byte(c))
	fmt.Println(len(c))

	for _, j := range c {
		fmt.Println(string(j))
	}

	fmt.Println("h" + "w")
}
