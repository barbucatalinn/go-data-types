package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var r32 float32 = 4
	var i32 float32 = 3

	a := complex(r32, i32)
	fmt.Printf("a's size is %d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

	var b complex64
	b = 4 + 5i
	fmt.Printf("d's size is %d bytes\n", unsafe.Sizeof(b))
	fmt.Printf("d's type is %s\n", reflect.TypeOf(b))

	var r64 float64 = 4
	var i64 float64 = 3

	c := complex(r64, i64)
	fmt.Printf("c's size is %d bytes\n", unsafe.Sizeof(c))
	fmt.Printf("c's type is %s\n", reflect.TypeOf(c))

	d := 4 + 5i
	fmt.Printf("d's size is %d bytes\n", unsafe.Sizeof(d))
	fmt.Printf("d's type is %s\n", reflect.TypeOf(d))
}
