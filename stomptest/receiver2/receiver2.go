// stompngo demo

package main

import (
	"fmt" //
  "log"
	"net"
	"os"
	"runtime"
	"github.com/gmallard/stompngo"
	"strings"
	"sync"
	"time"
)

var printMsgs bool = true
var wg sync.WaitGroup
var qname = "/queue/stompngo.srpub"
var mq = 2
var host = "localhost"
var hap = host + ":"

var incrCtl sync.Mutex
var numRecv int

func recMessages(c *stompngo.Connection, q string) {

	var error error

	fmt.Printf("Start for q: %s\n", q)
	// Receive phase
	headers := stompngo.Headers{"destination", q,
		"id", q}
	fmt.Printf("qhdrs: %v\n", headers)
	sc, error := c.Subscribe(headers)
	if error != nil {
		// Handle error properly
		log.Fatalf("sub error: %v\n", error)
	}
	for input := range sc {
		inmsg := input.Message.BodyString()
		if printMsgs {
			fmt.Println("queue:", q, "Next Receive: ", inmsg)
		}
		time.Sleep(1e9 / 100) // Crudely simulate message processing
		incrCtl.Lock()
		numRecv++
		incrCtl.Unlock()
		if strings.HasPrefix(inmsg, "***EOF***") {
			fmt.Printf("goteof: %v %v\n", q, inmsg)
			break
		}
		if !strings.HasPrefix(inmsg, q) {
			log.Fatalf("bad prefix: %v, %v\n", q, inmsg)
		}
	}
	error = c.Unsubscribe(headers)
	if error != nil {
		// Handle error properly
		log.Fatalf("unsub error: %v\n", error)
	}
	wg.Done()
}

func main() {
	fmt.Println("Start...")

	// create a net.Connection, and pass that into Connectionect
	nc, error := net.Dial("tcp", hap+os.Getenv("STOMP_PORT"))
	if error != nil {
		log.Fatal(error)
	}
	// Connection
	ch := stompngo.Headers{"login", "getter", "passcode", "recv1234",
		"accept-version", "1.1", "host", host}
	c, error := stompngo.Connect(nc, ch)
	if error != nil {
		log.Fatal(error)
	}
	for i := 1; i <= mq; i++ {
		qn := fmt.Sprintf("%d", i)
		wg.Add(1)
		go recMessages(c, qname+qn)
	}
	wg.Wait()
	fmt.Printf("Num received: %d\n", numRecv)
	// Disconnect
	nh := stompngo.Headers{}
	error = c.Disconnect(nh)
	if error != nil {
		log.Fatalf("discerr %v\n", error)
	}
	fmt.Println("done nc.Close()")
	nc.Close()
	ngor := runtime.NumGoroutine()
	fmt.Printf("egor: %v\n", ngor)
	select {
	case v := <-c.MessageData:
		fmt.Printf("frame2: %v\n", v)
	default:
		fmt.Println("Nothing to show")
	}
	fmt.Println("End... mq:", mq)
}
