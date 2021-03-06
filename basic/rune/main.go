package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	r := 'a'

	fmt.Printf("Size: %d\n", unsafe.Sizeof(r))
	fmt.Printf("Type: %s\n", reflect.TypeOf(r))

	fmt.Printf("Unicode CodePoint: %U\n", r)
	fmt.Printf("Character: %c\n", r)

	s := "0b£"
	fmt.Printf("%U\n", []rune(s))
	fmt.Println([]rune(s))
}
