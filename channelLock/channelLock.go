/*
	An experiment with using channels as a replacement for a formal "lock".
	Seems to work, no corruption of the updated resource.
*/
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

var (
	sd      = make(chan struct{})
	req     = make(chan chan bool)
	wanted  = 721
	current = 0 // The resource to update
)

func runner1(r int) {
	fmt.Println(r, "runner1 starts")
	var lc = 0
	var q = false
	for {
		select {
		case _ = <-sd:
			q = true
		default:
		}
		if q {
			break
		}
		rp := make(chan bool)
		req <- rp // Send request
		_ = <-rp  // Wait for reply (OK)
		if current == wanted {
			doSleep(r)
			continue
		}
		fmt.Println(r, "runner1 run starts", lc)
		doSleep(r)
		if current < wanted {
			lc++
			current++
		} else {
			fmt.Println(r, "runner1 run skips")
		}
		fmt.Println(r, "runner1 run ends", lc)
	}
	fmt.Println(r, "runner1 ends", lc)
}

func runner2(r int) {
	fmt.Println(r, "runner2 starts")
	var lc = 0
	var q = false
	for {
		select {
		case _ = <-sd:
			q = true
		default:
		}
		if q {
			break
		}
		rp := make(chan bool)
		req <- rp // Send request
		_ = <-rp  // Wait for reply (OK)
		if current == wanted {
			doSleep(r)
			continue
		}
		fmt.Println(r, "runner2 run starts", lc)
		doSleep(r)
		if current < wanted {
			lc++
			current++
		} else {
			fmt.Println(r, "runner2 run skips")
		}
		fmt.Println(r, "runner2 run ends", lc)
	}
	fmt.Println(r, "runner2 ends", lc)
}

func runner3(r int) {
	fmt.Println(r, "runner3 starts")
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
		rp := make(chan bool)
		req <- rp // Send request
		_ = <-rp  // Wait for reply (OK)
		if current == wanted {
			doSleep(r)
			continue
		}
		fmt.Println(r, "runner3 run starts", lc)
		doSleep(r)
		if current < wanted {
			lc++
			current++
		} else {
			fmt.Println(r, "runner3 run skips")
		}
		fmt.Println(r, "runner3 run ends", lc)
	}
	fmt.Println(r, "runner3 ends", lc)
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
}
func main() {
	//
	fmt.Printf("%s\n", "hi there")
	go lm() // Start lock manager
	i := 1
	go runner1(i)     // First worker
	go runner2(i + 1) // Second worker
	go runner3(i + 2) // Third worker
	// time.Sleep(1000 * time.Millisecond)
	for {
		time.Sleep(50 * time.Millisecond)
		if current == wanted {
			break
		}
	}
	close(sd)
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("%s %d\n", "current at end", current)
	fmt.Printf("%s\n", "bye bye")
}

func doSleep(r int) {
	n := rand.Intn(750)
	ns := fmt.Sprintf("%d", n) + "ms"
	d, e := time.ParseDuration(ns)
	if e != nil {
		log.Fatalln("Bad Parse:", e)
	}
	fmt.Printf("Runner %d sleeps %v\n", r, d)
	time.Sleep(d) // do some work
}
