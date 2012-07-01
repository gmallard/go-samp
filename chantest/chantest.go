/*
A demonstration of blocking when using an unbuffered channel.
*/
package main

import (
	"fmt"
	"time"
)

func worker(c chan int) {
	fmt.Println("worker starts ...")
	var x int = -1
	var ok bool = false
	time.Sleep(5 * 1e9)
	select {
	case x = <-c:
		ok = true
	default:
		ok = false
	}
	fmt.Printf("received %d %t\n", x, ok)
}

//
// Tests to see if a channel will block.
//
func main() {
	fmt.Println("Start ....")
	fmt.Println("===========")
	//
	var tosend int = 1
	c := make(chan int)
	go func() {
		time.Sleep(5 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}()
	//
	fmt.Println("sending", tosend)
	c <- tosend
	fmt.Println("sent1", tosend)
	//
	fmt.Println("===========")
	tosend = 2
	unbuf := make(chan int)
	go worker(unbuf) //
	unbuf <- tosend
	fmt.Println("sent2", tosend)
	//
	fmt.Println("===========")
	tosend = 3
	buf := make(chan int, 1)
	go worker(buf) //
	var ok bool = false
	select {
	case buf <- tosend:
		ok = true
	default:
		ok = false
	}
	fmt.Println("sent3", tosend, ok)
	time.Sleep(6 * 1e9)
	//
	fmt.Println("End ....")
}
