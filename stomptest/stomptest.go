// First gostomp demo

package main

import (
	"fmt" //
  "stomp"
  "net"
)

func main() {
	fmt.Println("Start...")

  // create a net.Conn, and pass that into Connect
	nc, error := net.Dial("tcp", "localhost:61613")
	if error != nil {
		// Handle error properly
	}

  // Connect
	ch := stomp.Header{"login": "userid", "passcode": "abcd1234"}
	c, error := stomp.Connect(nc, ch)
	if error != nil {
		// Handle error properly
	}

  // Send
  eh := stomp.Header{"header_key": "header_value"} // Extra headers
	error = c.Send("/queue/gostomp/pub001", eh, "gostomp message #1")
	if error != nil {
		// Handle error properly
	}

  // Disconnect
  nh := stomp.Header{}
	error = c.Disconnect(nh)
	if error != nil {
		// Handle error properly
	}

	fmt.Println("End...")
}
