// First gostomp demo

package main

import (
	"fmt" //
	"os"
  "net"
	"runtime"
  "stomp"
	"strings"
	"sync"
	"time"
)

var printMsgs bool = true
var wg sync.WaitGroup
var	qname = "/queue/gostomp.srpub"
var	mq = 100
var hap = "localhost:"

var incrCtl sync.Mutex
var numRecv int

func recMessages(c *stomp.Conn, q string) {

	var error os.Error

	fmt.Printf("Start for q: %s\n", q)

	// Receive phase
  headers := stomp.Header{"destination": q}
	_, error = c.Subscribe(headers)
	if error != nil {
		// Handle error properly
		fmt.Printf("sub error: %v\n", error)
	}
	for input := range c.Stompdata {
    inmsg := string(input.Message.Data)
    if printMsgs {
  		fmt.Println("queue:", q, "Next Receive: ", inmsg)
    }
		incrCtl.Lock()
		numRecv++
		incrCtl.Unlock()
		if strings.HasPrefix(inmsg, "***EOF***") {
			fmt.Printf("goteof: %v %v\n", q, inmsg)
			break
		}
	}

	wg.Done()
}

func main() {
	fmt.Println("Start...")

  // create a net.Conn, and pass that into Connect
	nc, error := net.Dial("tcp", hap + os.Getenv("STOMP_PORT"))
	if error != nil {
		// Handle error properly
	}

  // Connect
	ch := stomp.Header{"login": "getter", "passcode": "recv1234"}
	c, error := stomp.Connect(nc, ch)
	if error != nil {
		// Handle error properly
	}

	for i := 1; i <= mq; i++ {
		qn := fmt.Sprintf("%d", i)
		wg.Add(1)
		go recMessages(c, qname + qn)
	}
	wg.Wait()

	fmt.Println("unsubs: starts")

	for i := 1; i <= mq; i++ {
		qn := fmt.Sprintf("%d", i)
		error = c.Unsubscribe(stomp.Header{"destination": qn})
		if error != nil {
			// Handle error properly
			fmt.Printf("unsub error: %v\n", error)
		}
	}

	fmt.Printf("Num received: %d\n", numRecv)

  // Disconnect
  nh := stomp.Header{}
	error = c.Disconnect(nh)
	if error != nil {
		fmt.Printf("discerr %v\n", error)
	}

	nc.Close()

	time.Sleep(1e9 / 10)	// 100 ms

	ngor := runtime.Goroutines()
	fmt.Printf("egor: %v\n", ngor)
	if ngor > 1 {
		panic("too many gor")
	}

	fmt.Println("End...")
}
