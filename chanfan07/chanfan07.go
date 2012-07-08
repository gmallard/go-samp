/*
Example code based on Google I/O 2012 concurrency presentation.
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan string)
	c := boring("Joe!", quit)
	// Joe can olny talk so much, and then is told to quit.
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye!"
	fmt.Printf("Joe says: %s\n", <-quit)
}

func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() { // Launch from inside function
		i := 1
		for {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				i++
			case <-quit:
				fmt.Println("Cleanup complete")
				quit <- "See ya!"
				return
			}
		}
	}()
	return c // Return the channel to the caller
}
