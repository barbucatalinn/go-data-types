package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a bool

	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
	fmt.Printf("a's value is %t\n", a)

	// AND operation
	andOperation := 1 < 2 && 1 > 3
	fmt.Printf("Ouput of AND operation: %t\n", andOperation)

	// OR operation
	orOperation := 1 < 2 || 1 > 3
	fmt.Printf("Ouput of OR operation: %t\n", orOperation)

	// Negation operation
	negationOperation := !(1 > 2)
	fmt.Printf("Ouput of NEGATION operation: %t\n", negationOperation)
}
