/*
Example of using an unbuffered channel, based on gocourseday3.pdf.
*/
package main

//
import (
	"fmt"
	"time"
)

//
// gocourseday3.pdf - simple synchronous channel example.
// Slightly embellished.
//
func main() {
	fmt.Println("Start ....")
	//
	syn_chan := make(chan int)   // Synchronous: no buffering specified
	done_chan := make(chan bool) // Also synchronous
	go func() {
		time.Sleep(2 * 1e9)
		fmt.Println("receiver is up")
		x := <-syn_chan // Receive.  Will block until data received
		fmt.Println("starting work, received", x)
		tw := time.Duration(x * 1e9)
		time.Sleep(tw) // Simulate work
		fmt.Println("work complete")
		done_chan <- true // Signal done
	}()
	//
	fmt.Println("sending", 5)
	syn_chan <- 5 // Send.  Will block until a receive complete
	fmt.Println("sent", 5)
	fmt.Println("waiting for complete work")
	done_flag := <-done_chan
	fmt.Println("done is:", done_flag)
	//
	fmt.Println("End ....")
}
