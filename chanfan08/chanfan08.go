/*
Example code based on Google I/O 2012 concurrency presentation.
*/
package main

import (
	"fmt"
)

// Daisy chain

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	const n = 10000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}
