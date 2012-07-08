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
	//
	for {
		// Timeout each message 
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("You're too slow, I'm leaving")
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
