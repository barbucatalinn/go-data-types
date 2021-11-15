package main

import "fmt"

func main() {
	var a [3]string
	fmt.Println(fmt.Sprintf("%#v", a))

	a = [3]string{"a", "b"}

	fmt.Println(a)

	a[2] = "c"

	fmt.Println(a)
}
