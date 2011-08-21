// First gostomp demo

package main

import (
	"fmt" //
  "stomp"
  "net"
)

func main() {
	fmt.Println("Start...")

  var printMsgs bool = true
	qname := "/queue/gostomp.srpub"

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

	// Receive phase
  headers := make(stomp.Header) // empty headers
	_, error = c.Subscribe(qname, headers)
	if error != nil {
		// Handle error properly
	}
	for input := range c.Stompdata {
    inmsg := string(input.Message.Data)
    if printMsgs {
  		fmt.Println("Next Receive: ", inmsg)
    }
		if inmsg == "***EOF***" {
			break
		}
	}

  // Disconnect
  nh := stomp.Header{}
	error = c.Disconnect(nh)
	if error != nil {
		// Handle error properly
	}

	fmt.Println("End...")
}
