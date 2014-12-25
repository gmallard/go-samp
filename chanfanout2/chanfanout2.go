package main

import (
	"fmt"
	"time"
)

// Modified example from:
// http://stackoverflow.com/questions/16930251/go-one-producer-many-consumers
// and the referenced implementation.

func producer(iters int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < iters; i++ {
			c <- i
			time.Sleep(1 * time.Second)
		}
		close(c)
	}()
	return c
}

func consumer(cin <-chan int) {
	for i := range cin {
		fmt.Println(i)
	}
}

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

func fanOutUnbuffered(ch <-chan int, size int) []chan int {
	cs := make([]chan int, size)
	for i, _ := range cs {
		// The size of the channels buffer controls how far behind the recievers
		// of the fanOut channels can lag the other channels.
		cs[i] = make(chan int)
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
	c := producer(10)
	chans := fanOutUnbuffered(c, 3)
	go consumer(chans[0])
	go consumer(chans[1])
	consumer(chans[2])
}
