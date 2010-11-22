package main

import (
	"fmt"
  "time"
)

func testchan(ac chan int) {
	go func() {
		time.Sleep(5 * 1e9)
		x, ok := <-ac
		fmt.Printf("received %d %t\n", x, ok)
	}();
	//
	val := 10
	fmt.Println("sending", val)
	ok := ac <- val
	fmt.Printf("sent %d %t\n", val, ok)
	time.Sleep(10 * 1e9)
}
//
// Tests to see if a channel will block.
//
func main() {
  fmt.Println("Start ....")
	fmt.Println("===========")
	//
	c := make(chan int)
	go func() {
		time.Sleep(5*1e9)
		x := <-c
		fmt.Println("received", x)
	}();
	//
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
	//
	fmt.Println("===========")
	unbuf := make(chan int)
	testchan(unbuf) // Behavior seems ??? odd ...
	//
	fmt.Println("===========")
	buf := make(chan int, 1)
	testchan(buf)
	//
  fmt.Println("End ....")
}

