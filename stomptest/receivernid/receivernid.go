// stompngo demo

package main

import (
	"fmt" //
  "log"
	"net"
	"os"
	"runtime"
	"stomp"
	"strings"
	"sync"
	"time"
)

var printMsgs bool = true
var printHdrs bool = true
var wg sync.WaitGroup
var qname = "/queue/stompngo.srpub"
var mq = 2
var host = "localhost"
var hap = host + ":"

var incrCtl sync.Mutex
var numRecv int

func recMessages(c *stomp.Connection, q string) {

	var error error

	fmt.Printf("Start for q: %s\n", q)

	// Receive phase
	headers := stomp.Headers{"destination", q} // no ID here.  1.1 library should provide
	fmt.Printf("qhdrs: %v\n", headers)
	sc, error := c.Subscribe(headers)
	if error != nil {
		// Handle error properly
		fmt.Printf("sub error: %v\n", error)
	}
	first := true
	firstSub := ""
	for input := range sc {
		inmsg := string(input.Message.Body)
		if printHdrs {
			fmt.Println("queue:", q, "Next Receive: ", input.Message.Headers)
		}
		if printMsgs {
			fmt.Println("queue:", q, "Next Receive: ", inmsg)
		}

		firstSub = input.Message.Headers.Value("subscription")
		if first {
			if firstSub == "" {
				panic("first subscription header is empty")
			}
			fmt.Println("queue:", q, "FirstSub: ", firstSub)
			first = false
		} else {
			if firstSub != input.Message.Headers.Value("subscription") {
				panic(firstSub + " / " + input.Message.Headers.Value("subscription"))
			}
		}
		time.Sleep(1e9 / 100) // Crudely simulate message processing
		incrCtl.Lock()
		numRecv++
		incrCtl.Unlock()
		if strings.HasPrefix(inmsg, "***EOF***") {
			fmt.Println("queue:", q, "FirstSub:", firstSub, "goteof")
			break
		}
		if !strings.HasPrefix(inmsg, q) {
			fmt.Printf("bad prefix: %v, %v\n", q, inmsg)
			panic("bad prefix ....")
		}
		// Poll for adhoc errors
		select {
		case v := <-c.MessageData:
			fmt.Printf("frameError: %v\n", v.Message)
			fmt.Printf("frameError: [%v] [%v]\n", q, firstSub)
		default:
			fmt.Println("Nothing to show")
		}
	}
	uh := stomp.Headers{"id",firstSub,
    "destination", q}
	error = c.Unsubscribe(uh)
	if error != nil {
		log.Fatalf("unsub error: %v\n", error)
	}
	wg.Done()
}

func main() {
	fmt.Println("Start...")

	// create a net.Conn, and pass that into Connect
	nc, error := net.Dial("tcp", hap+os.Getenv("STOMP_PORT"))
	if error != nil {
		log.Fatal(error)
	}
	// Connect
	ch := stomp.Headers{"login", "getter", "passcode", "recv1234"}
	ch = ch.Add("accept-version","1.1")
	ch = ch.Add("host",host)
	c, error := stomp.Connect(nc, ch)
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
	nh := stomp.Headers{}
	error = c.Disconnect(nh)
	if error != nil {
		log.Fatalf("discerr %v\n", error)
	}
	fmt.Println("done nc.Close()")
	nc.Close()
	ngor := runtime.Goroutines()
	fmt.Printf("egor: %v\n", ngor)
	select {
	case v := <-c.MessageData:
		fmt.Printf("frame2: %s\n", v.Message.Command)
		fmt.Printf("header2: %v\n", v.Message.Headers)
		fmt.Printf("data2: %s\n", string(v.Message.Body))
	default:
		fmt.Println("Nothing to show")
	}
	fmt.Println("End... mq:", mq)
}
