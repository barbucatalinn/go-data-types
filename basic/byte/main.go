package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var r byte = 'a'

	fmt.Printf("Size: %d\n", unsafe.Sizeof(r))
	fmt.Printf("Type: %s\n", reflect.TypeOf(r))

	fmt.Printf("Character: %c\n", r)
	s := "abc"

	fmt.Println([]byte(s))
}
