/*

*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("starting")
	//
	c := make(chan int, 1) // A small buffer
	fmt.Printf("Print Test: %v\n", c)
	fmt.Printf("Length01: %d\n", len(c)) // Ignore races here :-)
	c <- 1
	fmt.Printf("Length02: %d\n", len(c)) // And here
	_ = <- c // clear

	// Test the if comma ok way
	c <- 2
	close(c) // the test
	// This receive works, even with the closed channel
	val, ok := <- c
	if ok {
		fmt.Printf("OK2, len, val: %d , %d\n", len(c), val) //
	} else {
		fmt.Printf("NOTOK2, len, val: %d , %d\n", len(c), val) //
	}
	// This receive fails
	val, ok = <- c
	if ok {
		fmt.Printf("OK3, len, val: %d , %d\n", len(c), val) //
	} else {
		fmt.Printf("NOTOK3, len, val: %d , %d\n", len(c), val) //
	}
	//
	// close(c) // 2nd time: panic


	// Test the range way
	c2 := make(chan int, 1) // A small buffer
	fmt.Printf("Print Test2: %v\n", c2)
	c2 <- 3
	close(c2) // the test
	//
	for v := range c2 {
		fmt.Printf("RANGE1, v: %d\n", v) //
	}
	fmt.Println("=============")
	for v2 := range c2 {
		fmt.Printf("RANGE2, v2: %d\n", v2) //  Does not run
	}
	//
	fmt.Println("ending")

	// Test the if comma ok way, no close
	c3 := make(chan int, 1) // A small buffer
	fmt.Printf("Print Test3: %v\n", c3)
	c3 <- 4
	// no close() here
	// This receive works, of course.
	val, ok = <- c3
	if ok {
		fmt.Printf("OK4, len, val: %d , %d\n", len(c3), val) //
	} else {
		fmt.Printf("NOTOK4, len, val: %d , %d\n", len(c3), val) //
	}
	// This panics
/*
	val4, ok4 := <- c3 //  -> fatal error: all goroutines are asleep - deadlock!
	if ok4 {
		fmt.Printf("OK4, len, val: %d , %d\n", len(c3), val4) //
	} else {
		fmt.Printf("NOTOK4, len, val: %d , %d\n", len(c3), val4) //
	}
*/
	// Try this instead.  Also panics.
/*
	fmt.Println("=============")
	for v3 := range c3 { //  -> fatal error: all goroutines are asleep - deadlock!
		fmt.Printf("RANGE3, v3: %d\n", v3) //
	}
*/

}

