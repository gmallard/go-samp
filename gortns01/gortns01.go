/*
Using sleep, and example from gocourseday3.pdf. 
*/
package main

//
import (
	"fmt"
	"time"
)

//
// The simplest example from gocourseday3.pdf, modified to prevent a 
// premature end to the example.
//
func IsReady(what string, minutes int64) {
	time.Sleep(time.Duration(minutes * 60 * 1e9))
	fmt.Println(what, "is ready")
}

func main() {
	fmt.Println("Starting .....")
	//
	go IsReady("tea", 6)
	go IsReady("coffee", 2)
	fmt.Println("I'm waiting....")

	// Crude synchronize with the go routines.
	time.Sleep(time.Duration(6.25 * 60 * 1e9))

	//
	fmt.Println("Ending .....")
}
