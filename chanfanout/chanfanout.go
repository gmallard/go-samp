package main

import (
	"fmt"
)

// Modified example from:
// http://stackoverflow.com/questions/16930251/go-one-producer-many-consumers

func fanOut(ch <-chan int, size, lag int) []chan int {
	cs := make([]chan int, size)
	for i, _ := range cs {
		// The size of the channels buffer controls how far behind the recievers
		// of the fanOut channels can lag the other channels.
		cs[i] = make(chan int, lag)
	}
	go func() {
		for i := range ch {
			for _, c := range cs {
				c <- i
			}
		}
		for _, c := range cs {
			// close all our fanOut channels when the input channel is exhausted.
			close(c)
		}
	}()
	return cs
}

func main() {
	fmt.Println("hi")
	const (
		chsz = 1
		sz   = 1
		lag  = 1
	)
	ch := make(chan int, chsz)
	co := fanOut(ch, sz, lag)
	fmt.Println("CO", co)
	fmt.Println("bye")
}
