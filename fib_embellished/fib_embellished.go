package main

import "fmt"

/*
Demonstrate:

a) multiple return values
b) use of := in functions

*/
func fib(i int) (result int, flag bool) {
	if i < 0 { return i, false }
	result, flag = i, true
	if i <= 1 { return }
	//
	resm1, flagm1 := fib(i - 1)
	resm2, flagm2 := fib(i - 2)
	return resm1 + resm2, flagm1 || flagm2
}

func main() {
	fmt.Printf("n \tfib(n)\tCheck\n")
	fmt.Printf("==\t======\t=====\n")
	//
	for ni := -10; ni < 21; ni++ {
		result, flag := fib(ni)
		if flag {
			fmt.Printf("%d\t%d\ttrue\n", ni, result)
		} else {
			fmt.Printf("%d\t%d\tfalse\n", ni, result)
		}
	}
}

