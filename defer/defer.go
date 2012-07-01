/*
A contrived example of using defer.
*/
package main

import (
	"fmt" //
)

/*
Defers operate LIFO
*/
func printBack() {
	for i := 1; i <= 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

/*
Using defer for tracing.
*/
func trace(s string) string {
	fmt.Println("entering: ", s)
	return s
}

//
func un(s string) {
	fmt.Println("leaving: ", s)
}
func inner() {
	defer un(trace("inner"))
	fmt.Println("in inner()")
}
func outer() {
	defer un(trace("outer"))
	fmt.Println("in outer() before call")
	inner()
	fmt.Println("in outer() after call")
}

func main() {
	fmt.Println("Start...")
	printBack()
	//
	fmt.Println()
	outer()
	//
	fmt.Println("End...")
}
