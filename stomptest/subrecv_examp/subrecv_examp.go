//

package main

import (
	"fmt" //
  "net"
  "os"
//  "strconv"
	"stomp"
//	"sync"
)

var nmsgs = 100
var	qname = "/queue/gostomp.subrecv_examp"
var host = "localhost"
var hap = host + ":"

func main() {

  // create a net.Conn, and pass that into Connect
	nc, error := net.Dial("tcp", hap + os.Getenv("STOMP_PORT"))
	if error != nil {
		// Handle error properly
	}

  // Connect
	ch := stomp.Header{"login": "putter", "passcode": "send1234"}

	//
	ch["accept-version"] = "1.1"
	ch["host"] = host

	c, error := stomp.Connect(nc, ch)
	if error != nil {
		panic(error)
	}

	sh := stomp.Header{"destination": qname}
	for i := 1; i <= nmsgs; i++ {
		msg := "subrecv message " + fmt.Sprintf("%d", i)
		error = c.Send(sh, msg)
		if error != nil {
			panic(error)
		}
		fmt.Printf("Sent message: %s\n", msg)
	}

	sc, error := c.Subscribe(sh)
	if error != nil {
		panic(error)
	}

	subid := ""
	i := 1
	for {
		// Sanity check.  Any unanticipated ERROR frames?
		select {
			case v := <- c.Stompdata:
				fmt.Printf("frame2: %s\n", v.Message.MsgFrame)
				fmt.Printf("header2: %v\n", v.Message.Header)
				fmt.Printf("data2: %s\n", string(v.Message.Data))
			default:
				fmt.Println("Nothing to show - 1")
		}
		fmt.Println("Start receive ....")
		d := <- sc
		fmt.Printf("d: %v\n", d)
		if d.Error != nil {
			panic(d.Error)
		}
		if i == 1 {
			subid = d.Message.Header["subscription"]
			fmt.Printf("Subscription is: %s\n", subid)
		}
		//
		fmt.Printf("Received message: %s\n", string(d.Message.Data))
		i++
		if i > nmsgs {
			break
		}
	}

	uh := stomp.Header{"destination": qname, "id": subid}
	error = c.Unsubscribe(uh)
	if error != nil {
		panic(error)
	}

  // Disconnect
  nh := stomp.Header{}
	error = c.Disconnect(nh)
	if error != nil {
		panic(error)
	}

	// Sanity check.  Any unanticipated ERROR frames?
	select {
		case v := <- c.Stompdata:
			fmt.Printf("frame2: %s\n", v.Message.MsgFrame)
			fmt.Printf("header2: %v\n", v.Message.Header)
			fmt.Printf("data2: %s\n", string(v.Message.Data))
		default:
			fmt.Println("Nothing to show - 2")
	}

	fmt.Println("done disconnect, start nc.Close()")
	nc.Close()
	fmt.Println("done nc.Close()")

}

