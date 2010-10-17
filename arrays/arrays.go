package main

import "fmt"
/*
From examples.
*/
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
// Function parameter: a pointer to an array of integers.
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
	var ar [3] int
	f(ar)
	fp(&ar)
	//
	// Array Literal
	//
	var arb = [3]int{ 1, 2, 3 }
	f(arb)
	fp(&arb)
	//
	// Array Literal
	//
	var arc = []int{ 1, 2, 3, 4, 5 }
	fa(arc)			// Note, different function needed (signature)
	fap(&arc)		// ....
	//
	// Array Literal, implied length of 5.
	//
	// This fails to compile:
	//
	// var ard = [...]int{ 1, 3, 5, 7, 9 }
	// fa(ard)
	// fap(&ard)
	//
	// with error messages:
	// ..... cannot use ard (type [5]int) as type []int in function argument
	// ..... cannot use &ard (type *[5]int) as type *[]int in function argument
	//
}

