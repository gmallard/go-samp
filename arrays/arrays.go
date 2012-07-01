/*
A demonstration of using go arrays and slices.
*/
package main

import "fmt"

//
// Function parameter: an array of 3 integers.
//
func f(a [3]int) {
	fmt.Println(a)
}

//
// Function parameter: an array of integers, any length
//
func fa(a []int) {
	fmt.Println(a)
}

//
// Function parameter: a pointer to an array of 3 integers.
//
func fp(a *[3]int) {
	fmt.Println(a)
}

//
// Function parameter: a pointer to an slice of integers.
//
func fap(a *[]int) {
	fmt.Println(a)
}

/*

*/
func main() {
	//
	// Declared array
	//
	var ar [3]int
	f(ar)
	fp(&ar)
	//
	// Array Literal
	//
	var arb = [3]int{1, 2, 3}
	f(arb)
	fp(&arb)
	//
	// This is actually/technically a slice literal, not an array literal.
	//
	var arc = []int{1, 2, 3, 4, 5}
	fa(arc)   // Note, different function needed (signature)
	fap(&arc) // ....
	//
	// ------------------------------------------------------------------------
	//
	// Array Literal, implied length of 5.
	//
	var ard = [...]int{1, 3, 5, 7, 9}
	//
	// This fails to compile:
	//
	// fa(ard)
	//
	// with error messages:
	// ..... cannot use ard (type [5]int) as type []int in function argument
	//
	// This fails to compile:
	//
	// fap(&ard)
	//
	// with error messages:
	// ..... cannot use &ard (type *[5]int) as type *[]int in function argument
	//
	// ------------------------------------------------------------------------
	fmt.Printf("%d\n", len(ard))
	//
	// Loop through array - one
	//
	for i := 0; i < len(arb); i++ {
		fmt.Printf("%d\n", arb[i])
	}
	//
	// Loop through array - two
	//
	for _, narb := range arb {
		fmt.Printf("%d\n", narb)
	}
}
