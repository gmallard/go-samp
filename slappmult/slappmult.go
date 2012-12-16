/*
Demonstrate append multiple slices.
*/
package main

import (
	"fmt"
)

func addmult(orig []byte, stuff ...[]byte) []byte {
	b := make([]byte, 0)
	b = append(b, orig...)
	for _, val := range stuff {
		b = append(b, val...)
	}
	return b
}

func addmultv(stuff [][]byte) []byte {
	b := make([]byte, 0)
	for _, val := range stuff {
		b = append(b, val...)
	}
	return b
}

// Append multiple splices.
func main() {
	fmt.Println("Hi")
	s1 := "1\n"
	s1b := []byte(s1)
	s2 := "2\n"
	s2b := []byte(s2)
	s3 := "3\n"
	s3b := []byte(s3)

	//
	fmt.Printf("%v", s1)
	fmt.Printf("%v\n", s1b)
	fmt.Printf("%v", s2)
	fmt.Printf("%v\n", s2b)
	fmt.Printf("%v", s3)
	fmt.Printf("%v\n", s3b)

	// Here is one way to do this.
	ob := make([]byte, 0)
	ob = append(ob, s1b...)
	ob = append(ob, s2b...)
	ob = append(ob, s3b...)
	fmt.Printf("%v\n", ob)

	// And here is a somewhat more generalized way.
	ob2 := make([]byte, 0) // An empty orig .....
	ob2 = addmult(ob2, s1b, s2b, s3b)
	fmt.Printf("%v\n", ob2)

	// And here is yet another somewhat more generalized way.
	bat := addmultv([][]byte{s1b, s2b, s3b})
	fmt.Printf("%v\n", bat)

	// Is there a way to do this with a single append call ???
	// Is there a better way to do this in general ???

}
