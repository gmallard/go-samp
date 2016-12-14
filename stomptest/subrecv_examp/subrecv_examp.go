/*
Send and receive STOMP messages using https://github.com/gmallard/stompngo using
a STOMP 1.1 broker.
*/
package main

import (
	"fmt" //
	"log"
	"net"
	"os"

	"github.com/gmallard/stompngo"
)

var nmsgs = 10
var qname = "/queue/stompngo.subrecv_examp"
var host = "localhost"
var hap = host + ":"

func main() {

	// create a net.Conn, and pass that into Connect
	p := os.Getenv("STOMP_PORT")
	if p == "" {
		p = "61613"
	}
	nc, error := net.Dial("tcp", hap+p)
	if error != nil {
		// Handle error properly
	}
	// Connect
	ch := stompngo.Headers{"login", "guest", "passcode", "guest",
		"accept-version", "1.1", "host", host}
	c, error := stompngo.Connect(nc, ch)
	if error != nil {
		log.Fatal(error)
	}
	sh := stompngo.Headers{"destination", qname}
	for i := 1; i <= nmsgs; i++ {
		msg := "subrecv message " + fmt.Sprintf("%d", i)
		error = c.Send(sh, msg)
		if error != nil {
			log.Fatal(error)
		}
		fmt.Printf("Sent message: %s\n", msg)
	}
	// No 'id' header is present -> the stompngo client library creates one.
	// The assigned subscription id must be determined from subsequent traffic.
	sc, error := c.Subscribe(sh)
	if error != nil {
		log.Fatal(error)
	}
	subid := ""
	i := 1
	for {
		// Sanity check.  Any unanticipated ERROR frames?
		select {
		case v := <-c.MessageData:
			log.Fatalf("frame1: %v\n", v)
		default:
			fmt.Println("Nothing to show - 1")
		}
		fmt.Println("Start receive ....")
		d := <-sc
		// fmt.Printf("d: %v\n", d)
		if d.Error != nil {
			log.Fatal(d.Error)
		}
		// Save the subscription id in use for the eventual UNSUBSCRIBE.
		if i == 1 {
			subid = d.Message.Headers.Value("subscription")
			fmt.Printf("Subscription is: %s\n", subid)
		}
		//
		fmt.Printf("Received message: %s\n", d.Message.BodyString())
		i++
		if i > nmsgs {
			break
		}
	}
	uh := stompngo.Headers{"destination", qname, "id", subid} // Unsubscribe headers
	error = c.Unsubscribe(uh)
	if error != nil {
		log.Fatal(error)
	}
	// Disconnect
	nh := stompngo.Headers{}
	error = c.Disconnect(nh)
	if error != nil {
		log.Fatal(error)
	}
	// Sanity check.  Any unanticipated ERROR frames?
	select {
	case v := <-c.MessageData:
		log.Fatalf("frame2: %v\n", v)
	default:
		fmt.Println("Nothing to show - 2")
	}
	fmt.Println("done disconnect, start nc.Close()")
	nc.Close()
	fmt.Println("done nc.Close()")

}
