/*
Using type coercion with automatic declares.
*/
package main

import (
	"fmt" //
)

func main() {
	fmt.Println("Start...")
	// Corece type with auto declaration
	i := int32(0)
	fmt.Printf("%v\n", i)
	// Corece type with auto declaration
	j := int64(5)
	fmt.Printf("%v\n", j)
	fmt.Println("End...")
}
