/*
Loop and receive values from a buffered channel.
*/
package main

//
import (
	"fmt"
	//	"time"
)

//
func runner(inc chan int, wc chan bool) {
	fmt.Println("runner starting")
	dobreak := false
	for {
		if dobreak {
			break
		}
		select {
		case ival := <-inc:
			fmt.Println("runner got value", ival)
		default:
			fmt.Println("runner found nothing")
			dobreak = true
		}
	}
	//
	wc <- true // we are done
	fmt.Println("runner ending")
}

//
func main() {
	fmt.Println("Start ....")
	//
	isc := make(chan int, 3)
	isc <- 1
	isc <- 2
	isc <- 3
	//
	fmt.Println("Three are queued ....")
	waitChan := make(chan bool)
	go runner(isc, waitChan)
	fmt.Println("Starting wait ....")
	doneFlag := <-waitChan
	fmt.Println("Done", doneFlag)
	//
	fmt.Println("End ....")
}
