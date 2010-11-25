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
	if i < 0 {
		return i, false
	}
	result, flag = i, true
	if i <= 1 {
		return
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
	for ni := -10; ni < 21; ni++ {
		result, flag := fib(ni)
		fmt.Printf("%d\t%d\t%t\n", ni, result, flag)
	}
}
