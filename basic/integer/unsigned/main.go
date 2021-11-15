package main

import (
	"fmt"
	"math/bits"
	"reflect"
	"unsafe"
)

func main() {
	// computed as const uintSize = 32 << (^uuint(0) >> 32 & 1) // 32 or 64
	sizeOfUintInBits := bits.UintSize
	fmt.Printf("%d bits\n", sizeOfUintInBits)

	var a uint
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

	var b uint = 2
	fmt.Printf("b's typs is %s\n", reflect.TypeOf(b))

	// uint8
	var c uint8 = 2
	fmt.Printf("%d bytes\n", unsafe.Sizeof(c))
	fmt.Printf("c's type is %s\n", reflect.TypeOf(c))

	// uint16
	var d uint16 = 2
	fmt.Printf("%d bytes\n", unsafe.Sizeof(d))
	fmt.Printf("d's type is %s\n", reflect.TypeOf(d))

	// uint32
	var e uint32 = 2
	fmt.Printf("%d bytes\n", unsafe.Sizeof(e))
	fmt.Printf("e's type is %s\n", reflect.TypeOf(e))

	// uint64
	var f uint64 = 2
	fmt.Printf("%d bytes\n", unsafe.Sizeof(f))
	fmt.Printf("f's type is %s\n", reflect.TypeOf(f))
}
