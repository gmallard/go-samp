package main

import (
	"fmt"
	"container/vector"
)

func procf(af interface {}) {
	fmt.Printf("ProcN: %g\n", af)
}

func main() {
  fmt.Println("Start ...")
	//
	var v = new(vector.Vector)
	var f float = 1234.5
	//
	v.Push(f)
	fmt.Printf("Len: %d\n", v.Len())
	fmt.Printf("Cap: %d\n", v.Cap())
	//
	x := v.At(0)
	fmt.Printf("ELT0: %g\n", x)
	//
	f = 2345.6
	v.Push(f)
	fmt.Printf("Len: %d\n", v.Len())
	fmt.Printf("Cap: %d\n", v.Cap())
	y := v.At(1)
	fmt.Printf("ELT1: %g\n", y)
	//
	for i := 0; i < v.Len(); i++ {
		fmt.Printf("ELTn: %g\n", v.At(i))
	}
	//
	v.Do(procf)	// Call for all elements
	//
	if y != 1234. {} // OK, compiles ( compile-time err in PDF ?? )
	if y.(float) != 1234. {} // OK, compiles
	// if y.(int) != 1234 {} // run-time err ->
	// panic: interface conversion: interface is float, not int
	//
  fmt.Println("Bye ...")
}

