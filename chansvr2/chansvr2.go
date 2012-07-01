/*
A simple server using channels, with client controlled shutdown.
*/
package main

import "fmt"

//
// Another example from the gocourseday3.pdf.
//
// This example works as presented in the PDF.  This is mostly cut and 
// paste.  This adds to the previous example, but shows client controlled
// shutdown of the server.
//
// The request.
//
type request struct {
	a, b   int
	replyc chan int
}

//
// Operation type definition.
//
type binOp func(a, b int) int

//
// The low level 'runner'.
//
func run(op binOp, req *request) {
	req.replyc <- op(req.a, req.b)
}

//
// The server logic.
//
func server(op binOp, service chan *request, quit chan bool) {
	for {
		select {
		case req := <-service:
			go run(op, req) // don't wait for it
		case <-quit:
			fmt.Println("OK, serever done")
			return
		}
	}
}

//
// Server startup.
//
func startServer(op binOp) (chan *request, chan bool) {
	service := make(chan *request)
	quit := make(chan bool)
	go server(op, service, quit)
	return service, quit
}

//
// Print requests nicely.
//
func (r *request) String() string {
	return fmt.Sprintf("%d+%d=%d", r.a, r.b, <-r.replyc)
}

//
// Mainline.
//
func main() {
	fmt.Println("Start...")
	// Server start.
	var adderChan, quitChan = startServer(func(a, b int) int { return a + b })
	// Create some requests
	req1 := &request{7, 8, make(chan int)}
	req2 := &request{17, 18, make(chan int)}
	// Send the requests
	adderChan <- req1
	adderChan <- req2
	// Get and show reply.
	fmt.Println(req2, req1)
	//
	quitChan <- true
	//
	fmt.Println("End...")
}
