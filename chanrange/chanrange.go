/*
Loop and receive values from a buffered channel using range.
*/
package main

//
import (
	"fmt"
)

//
func runner(inc chan int, wc chan bool) {
	fmt.Println("runner starting")
	for ival := range inc {
		fmt.Println("runner got value", ival)
		// If we do not break here, we get at runtime:
		// throw: all goroutines are asleep - deadlock!
		if ival == 3 {
			break
		}
	}
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
