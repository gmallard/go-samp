/*
The obligatory Fibonacci number example.
*/
package main

// Import formatting package
import "fmt"

//
// A somewhat classical recursive function to determine the n'th
// fibonacci number.  This includes the definition of fibinacci 
// numbers for negative integers.
//
// Reference: http://en.wikipedia.org/wiki/Fibonacci_number
//
func fib(i int) int {
	// The basics:  0 and 1
	if i == 0 || i == 1 {
		return i
	}
	// For negative integers.
	if i < 0 {
		posfib := fib(-1 * i)
		// See formulae at the above reference.
		if (-1*i+1)%2 != 0 {
			return -posfib
		}
		return posfib
	}
	// And for all other positives.
	return fib(i-1) + fib(i-2)
}

//
// Mainline.
//
func main() {
	// List some fibonacci numbers.
	fmt.Printf("n \tfib(n)\n")
	fmt.Printf("==\t======\n")
	for x := -10; x < 11; x++ {
		fmt.Printf("%d\t%d\n", x, fib(x))
	}
}
