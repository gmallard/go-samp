/*
Originally meant to be a test of overloading, which of course is not implemented.
*/
package main

import "fmt"

/*
From examples.  Interesting notes:

a) Function overloads are unsupported :-(.
b) Arrays are values
c) Can take address of an Array

*/
func f(a [3]int) { fmt.Println(a) }

// func f(a *[3]int) { fmt.Println(a) } // fails to compile
func fp(a *[3]int) { fmt.Println(a) }

/*
 */
func main() {
	fmt.Printf("Basic overload tester ....\n")
	var ar [3]int
	f(ar)
	// passes a copy of ar
	fp(&ar) // passes a pointer to ar
}
