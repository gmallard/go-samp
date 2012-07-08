/*
Example code based on Google I/O 2012 concurrency presentation.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quit := make(chan bool)
	c := boring("Joe!", quit)
	// Joe can olny talk so much, and then is told to quit.
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
	time.Sleep(100 * time.Millisecond)
}

func boring(msg string, quit chan bool) <-chan string {
	c := make(chan string)
	go func() { // Launch from inside function
		i := 1
		for {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				i++
			case <-quit:
				fmt.Println("Stopping")
				return
			}
		}
	}()
	return c // Return the channel to the caller
}
