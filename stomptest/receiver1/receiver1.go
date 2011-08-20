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
  for i := 0; i < 10; i++ {
	  error = c.Send("/queue/gostomp/pub001", eh, "gostomp message #" + string(i))
	  if error != nil {
		  // Handle error properly
	  }
  }

	// Receive phase
	queue_name := "/queue/gostomp/pub001"
  headers := make(stomp.Header) // empty headers
	error = c.Subscribe(queue_name, headers)
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
