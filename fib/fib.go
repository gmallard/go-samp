package main

import "fmt"

func fib(i int) int {
	if i < 0 { return i }
	if i == 0 { return i }
	if i == 1 { return i }	
	return fib(i-1) + fib(i-2)
}

func main() {
  fmt.Printf("n \tfib(n)\n")
  fmt.Printf("==\t======\n")
	for x:= 0; x < 20; x++ {
		fmt.Printf("%d\t%d\n", x, fib(x))
	}
}

