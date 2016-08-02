package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sd  = make(chan struct{})
	req = make(chan chan bool)
	// rep  = make(chan bool)
	lock sync.RWMutex
)

func runner1() {
	fmt.Println("runner1 starts")
	var lc = 0
	var q = false
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
		lc++
		rp := make(chan bool)
		req <- rp // Send request
		_ = <-rp  // Wait for reply (OK)
		fmt.Println("runner1 run starts", lc)
		time.Sleep(500 * time.Millisecond) // do some work
		fmt.Println("runner1 run ends", lc)
	}
	fmt.Println("runner1 ends", lc)
}

func runner2() {
	fmt.Println("runner2 starts")
	var lc = 0
	var q = false
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
		lc++
		rp := make(chan bool)
		req <- rp // Send request
		_ = <-rp  // Wait for reply (OK)
		fmt.Println("runner2 run starts", lc)
		time.Sleep(500 * time.Millisecond) // do some work
		fmt.Println("runner2 run ends", lc)
	}
	fmt.Println("runner2 ends", lc)
}

func lm() {
	fmt.Println("lm starts")
	//_ = <-req
	var lc = 0
	var q = false
	var rp = make(chan bool)
	for {
		select {
		case _ = <-sd:
			q = true
		case rp = <-req:
		}
		if q {
			break
		}
		lc++
		fmt.Println("lm have a request", lc)
		rp <- true
		fmt.Println("lm sent reply", lc)
	}
	fmt.Println("lm ends", lc)
	//rep <- true
}
func main() {
	//
	fmt.Printf("%s\n", "hi there")
	go lm()      // Start lock manager
	go runner1() // First worker
	go runner2() // Second worker
	time.Sleep(1000 * time.Millisecond)
	// req <- true
	// _ = <-rep
	// sd <- true
	close(sd)
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("%s\n", "bye bye")
}
