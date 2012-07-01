/*
A demonstration of using a buffered channel.  Originally from gocourseday3.pdf.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start ....")

	c := make(chan int, 50)
	go func() {
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
	time.Sleep(10 * time.Second)
	//
	fmt.Println("End ....")
}
