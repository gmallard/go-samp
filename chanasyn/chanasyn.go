package main
//
import (
	"fmt"
	"time"
)
//
// gocourseday3.pdf - simple asynchronous channel example.
//
func main() {
  fmt.Println("Start ....")

	c := make(chan int, 50);
	go func() {
		time.Sleep(60*1e9);
		x := <-c;
		fmt.Println("received", x);
	}();
	fmt.Println("sending", 10);
	c <- 10;
	fmt.Println("sent", 10);
	time.Sleep(1.25 * 60 * 1e9)
	//
  fmt.Println("End ....")
}

