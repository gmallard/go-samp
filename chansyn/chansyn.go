package main
//
import (
	"fmt"
	"time"
)
//
// gocourseday3.pdf - simple synchronous channel example.
//
func main() {
	fmt.Println("Start ....")

	c := make(chan int)
	go func() {
		time.Sleep(60 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}()
	//
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
	//
	fmt.Println("End ....")
}
