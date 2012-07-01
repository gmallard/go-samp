/*
Using function literals.
*/
package main

import (
	"fmt" //
)

/*
Functions cannot be defined inside functions.
However, function literals may be defined and used.
*/
func runLit() {
	//
	g := func(i int) { fmt.Printf("%d\n", i) }
	//
	for i := 0; i < 10; i++ {
		g(i)
	}
}

/*
Function literals are closures.
*/
func adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

//
func main() {
	fmt.Println("Start...")
	//
	runLit()
	//
	fmt.Println()
	var f = adder()
	fmt.Println(f(1))
	fmt.Println(f(20))
	fmt.Println(f(300))
	//
	fmt.Println("End...")
}
