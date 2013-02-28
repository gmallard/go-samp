/*
An example taken from: https://groups.google.com/forum/?fromgroups#!topic/golang-nuts/VdmoZ59jjoE
*/
package main

import "fmt"

//
// See:
// https://groups.google.com/forum/?fromgroups#!topic/golang-nuts/VdmoZ59jjoE
// http://play.golang.org/p/ZJamJZgeEK
//
var (
	maxmsgs         = 10  // Anything, >= 1
	driver_out_init = 100 // Channel size
	driver_in_init  = 10  // Channel size
)

//
// A channel shunt (or shovel).  Example 2.
//
// Differences from example 1:
//
// This example changes variable names.  So that they make more sense to me.
// And also adds a number of comments.
//
// Parms:
// *input: a read only channel
// *output: a write only channel
func shunt(input <-chan int, output chan<- int) {
	var (
		i         int        // This is the channel payload, e.g.
		ok        bool       // Work var, for closed check .....
		shunt_in  = input    // Start by copying the input only channel
		shunt_out chan<- int // A local output only channel
	)
	for {
		select {

		// In practice, the following two 'cases' will flip-flop execution.
		// This happens because of:
		// a) the reassignment of the 'shunt_' channels *and*
		// b) the subtle re-assignment of the 'shunt_' input/output channels to 
		// 'nil', in an alternating sequence.
		//
		// Also note:
		// The 'first' case to ever be selected will be read from 'shunt_in'.

		case i, ok = <-shunt_in: // Read the input only channel
			if !ok { // Closed?
				close(output)       // Yep, close the external output channel
				fmt.Println("done") // Blah, ....
				return              // We are done
			}
			fmt.Println("receiver shunt in", i) // We have some input, print it

			// Flip local output to external output for next time
			shunt_out = output // Make sure we send properly next loop

			// Flip local input channel to never ready
			shunt_in = nil // Make sure we do not have input next loop

		case shunt_out <- i: // Write to the output only channel
			fmt.Println("sender shunt out", i) // Show results

			// Flip local input to external input for next time
			shunt_in = input // Make sure we receive properly next loop

			// Flip local output channel to never ready
			shunt_out = nil // Make sure we do no output next loop
		}
	}

}

func main() {

	driver_out := make(chan int, driver_out_init)
	go func() { // Send data to a channel, separate goroutine
		for i := 0; i < maxmsgs; i++ { // All of these times ...
			driver_out <- i              // Really send the data
			fmt.Println("sent asy: ", i) // Report
		}
		fmt.Println("asy sends done")     // Report
		close(driver_out)                 // Close, no more data to send
		fmt.Println(":-) Sends are done") // Report
	}()

	// Start receives as soon as possible.
	driver_in := make(chan int, driver_in_init) // Receiver for the coordinated results
	// In main, we acually just read these results below.

	go shunt(driver_out, driver_in) // Start the shunt: this really starts things ...

	// And receive all of the intermediate values.
	for i := range driver_in { // Read each and every next 'final' value
		fmt.Println("main driver_in read:", i) // Show it ....
	}
	fmt.Println("exit") // Bye ...

}
