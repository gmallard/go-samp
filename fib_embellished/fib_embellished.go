/*
An embellished Fibinacci number example.
*/
package main

//
import "fmt"

/*
Demonstrate:

a) multiple return values
b) use of := in functions

*/

//
//         -6 -5 -4 -3 -2 -1  0  1  2  3  4  5  6
//     ...,-8, 5,-3, 2,-1, 1, 0, 1, 1, 2, 3, 5, 8,...
//
// Reference: http://en.wikipedia.org/wiki/Fibonacci_number
//
func fib(i int) (result int, flag bool) {
	// The basics:  0 and 1
	if i == 0 || i == 1 {
		return i, true
	}
	// For negative integers.
	if i < 0 {
		posfib, ok := fib(-1 * i)
		// See formulae at the above reference.
		if (-1*i+1)%2 != 0 {
			return -posfib, ok
		}
		return posfib, ok
	}
	//
	resm1, flagm1 := fib(i - 1)
	resm2, flagm2 := fib(i - 2)
	return resm1 + resm2, flagm1 || flagm2
}

func main() {
	fmt.Printf("n \tfib(n)\tCheck\n")
	fmt.Printf("==\t======\t=====\n")
	//
	for ni := -7; ni <= 6; ni++ {
		result, flag := fib(ni)
		fmt.Printf("%d\t%d\t%t\n", ni, result, flag)
	}
}
