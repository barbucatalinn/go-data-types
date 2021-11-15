package main

import (
	"fmt"
	"math/bits"
	"reflect"
	"unsafe"
)

func main() {
	// computed as const uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	sizeOfIntInBits := bits.UintSize
	fmt.Printf("%d bits\n", sizeOfIntInBits)

	var a int
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

	b := 2
	fmt.Printf("b's typs is %s\n", reflect.TypeOf(b))

	// int8
	var c int8 = 2
	fmt.Printf("%d bytes\n", unsafe.Sizeof(c))
	fmt.Printf("c's type is %s\n", reflect.TypeOf(c))

	// int16
	var d int16 = 2
	fmt.Printf("%d bytes\n", unsafe.Sizeof(d))
	fmt.Printf("d's type is %s\n", reflect.TypeOf(d))

	// int32
	var e int32 = 2
	fmt.Printf("%d bytes\n", unsafe.Sizeof(e))
	fmt.Printf("e's type is %s\n", reflect.TypeOf(e))

	// int64
	var f int64 = 2
	fmt.Printf("%d bytes\n", unsafe.Sizeof(f))
	fmt.Printf("f's type is %s\n", reflect.TypeOf(f))
}
