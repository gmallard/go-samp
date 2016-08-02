package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sd   = make(chan struct{})
	req  = make(chan bool)
	rep  = make(chan bool)
	lock sync.RWMutex
)

func runner1() {
	fmt.Println("runner1 starts")
	q := false
	// r := true
	for {
		select {
		case _ = <-sd:
			q = true
		default:
		}
		if q {
			break
		}
		req <- true // Send request
		_ = <-rep   // Wait for reply (OK)
		fmt.Println("runner1 run starts")
		time.Sleep(500 * time.Millisecond) // do some work
		fmt.Println("runner1 run ends")
	}
	fmt.Println("runner1 ends")
}
func lm() {
	fmt.Println("lm starts")
	//_ = <-req
	q := false
	for {
		select {
		case _ = <-sd:
			q = true
		case _ = <-req:
		}
		if q {
			break
		}
		fmt.Println("lm have a request")
		rep <- true
		fmt.Println("lm sent reply")
	}
	fmt.Println("lm ends")
	//rep <- true
}
func main() {
	//
	fmt.Printf("%s\n", "hi there")
	go lm()
	go runner1()
	time.Sleep(250 * time.Millisecond)
	// req <- true
	// _ = <-rep
	// sd <- true
	close(sd)
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("%s\n", "bye bye")
}
