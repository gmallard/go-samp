/*
Yet another Fibonacci number example using a user defined operator.
*/
package main

import "fmt"

/*
User supplied add function.
*/
func myadd(i, j int) int {
	return i + j
}

/*
Demonstrate:

a) Use of user supplied add operation

*/
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
	return myadd(resm1, resm2), flagm1 || flagm2
}

func main() {
	fmt.Printf("n \tfib(n)\tCheck\n")
	fmt.Printf("==\t======\t=====\n")
	//
	for ni := -6; ni < 7; ni++ {
		result, flag := fib(ni)
		fmt.Printf("%d\t%d\t%t\n", ni, result, flag)
	}
}
