// First gostomp demo

package main

import (
	"fmt" //
  "net"
  "os"
	"runtime"
  "stomp"
  "strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup
var printMsgs bool = true
var  nmsgs = 5000
var	qname = "/queue/gostomp.srpub"
var	mq = 100
var host = "localhost"
var hap = host + ":"


func sendMessages(c *stomp.Conn, q string, n int, k int) {

	var error os.Error
//	ks := q + fmt.Sprintf("%d", k)

  // Send
  eh := stomp.Header{"destination": q} // Extra headers
  for i := 1; i <= n; i++ {
		m := q + " gostomp message #" + strconv.Itoa(i)
		if printMsgs {
			fmt.Println("msg:", m)
		}
	  error = c.Send(eh, m)
	  if error != nil {
		  fmt.Printf("send error: %v\n", error)
	  }
		//
		time.Sleep(1e9 / 100)	// Simulate message build
  }

  error = c.Send(eh, "***EOF*** " + q)
  if error != nil {
	  // Handle error properly
  }
	wg.Done()

}

func main() {
	fmt.Println("Start...")

  //
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

	for i := 1; i <= mq; i++ {
		qn := fmt.Sprintf("%d", i)
		wg.Add(1)
		go sendMessages(c, qname + qn, nmsgs, i)

	}
	wg.Wait()
	fmt.Println("done with wait")

  // Disconnect
  nh := stomp.Header{}
	error = c.Disconnect(nh)
	if error != nil {
		fmt.Printf("discerr %v\n", error)
	}

	fmt.Println("done disconnect, start nc.Close()")
	nc.Close()

	fmt.Println("done nc.Close()")
/*
	fmt.Println("start sleep")
	time.Sleep(1e9 / 10)	// 100 ms
	fmt.Println("end sleep")
*/

	ngor := runtime.Goroutines()
	fmt.Printf("egor: %v\n", ngor)

	select {
		case v := <- c.Stompdata:
			fmt.Printf("frame2: %s\n", v.Message.MsgFrame)
			fmt.Printf("header2: %v\n", v.Message.Header)
			fmt.Printf("data2: %s\n", string(v.Message.Data))
		default:
			fmt.Println("Nothing to show")
	}
/*
	if ngor > 1 {
		panic("too many gor")
	}
*/
	fmt.Println("End... ngor:", mq, " nmsgs:", nmsgs)
}
