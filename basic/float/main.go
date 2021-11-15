package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a float32 = 2

	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

	var b float64 = 2

	fmt.Printf("%d bytes\n", unsafe.Sizeof(b))
	fmt.Printf("b's type is %s\n", reflect.TypeOf(b))
}
