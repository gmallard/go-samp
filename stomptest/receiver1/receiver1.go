// gostompgo demo

package main

import (
	"fmt" //
	"net"
	"os"
	"runtime"
	"stomp"
	"strings"
	"sync"
	//	"time"
)

var printMsgs bool = true
var wg sync.WaitGroup
var qname = "/queue/gostomp.srpub"
var mq = 2
var hap = "localhost:"

var incrCtl sync.Mutex
var numRecv int

func recMessages(c *stomp.Connection, q string) {

	var error error

	fmt.Printf("Start for q: %s\n", q)

	// Receive phase
	headers := stomp.Headers{"destination", q}
	fmt.Printf("qhdrs: %v\n", headers)
	_, error = c.Subscribe(headers)
	if error != nil {
		// Handle error properly
		fmt.Printf("sub error: %v\n", error)
	}
	for input := range c.MessageData {
		inmsg := input.Message.BodyString()
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

	// create a net.Connection, and pass that into Connectionect
	nc, error := net.Dial("tcp", hap+os.Getenv("STOMP_PORT"))
	if error != nil {
		// Handle error properly
	}

	// Connect
	ch := stomp.Headers{"logiI121n", "getter", "passcode", "recv1234"}
	c, error := stomp.Connect(nc, ch)
	if error != nil {
		// Handle error properly
	}

	for i := 1; i <= mq; i++ {
		qn := fmt.Sprintf("%d", i)
		wg.Add(1)
		go recMessages(c, qname+qn)
	}
	wg.Wait()

	select {
	case v := <-c.MessageData:
		fmt.Printf("frame1: %v\n", v)
	default:
		fmt.Println("Nothing to show")
	}

	fmt.Println("unsubs: starts")

	for i := 1; i <= mq; i++ {
		qn := fmt.Sprintf("%d", i)
		h := stomp.Headers{"destination", qname + qn}
		fmt.Printf("unsubhdr: %v\n", h)
		error = c.Unsubscribe(h)
		if error != nil {
			// Handle error properly
			fmt.Printf("unsub error: %v\n", error)
		}
	}

	fmt.Printf("Num received: %d\n", numRecv)

	// Disconnect
	nh := stomp.Headers{}
	error = c.Disconnect(nh)
	if error != nil {
		fmt.Printf("discerr %v\n", error)
	}

	fmt.Println("done nc.Close()")
	nc.Close()

	/*
		fmt.Println("start sleep")
		time.Sleep(1e9 / 10)	// 100 ms
		fmt.Println("end sleep")
	*/

	ngor := runtime.Goroutines()
	fmt.Printf("egor: %v\n", ngor)

	select {
	case v := <-c.MessageData:
		fmt.Printf("frame2: %v\n", v)
	default:
		fmt.Println("Nothing to show")
	}
	/*
		if ngor > 1 {
			panic("too many gor")
		}
	*/
	fmt.Println("End...")
}
