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
	c := boring("Joe!")
	// Timeout entire loop
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much, I'm leaving")
			return
		}
	}
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() { // Launch from inside function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller
}
