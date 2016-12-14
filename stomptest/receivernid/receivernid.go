/*
Receive STOMP messages using https://github.com/gmallard/stompngo and a STOMP
1.0 broker, verify library added subscription ID.
*/
package main

import (
	"fmt" //
	"log"
	"net"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/gmallard/stompngo"
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

func recMessages(c *stompngo.Connection, q string) {

	var error error

	fmt.Printf("Start for q: %s\n", q)

	// Receive phase
	headers := stompngo.Headers{"destination", q} // no ID here.  1.1 library should provide
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
			fmt.Println("Headers for queue:", q, "Next Receive: ", input.Message.Headers)
		}
		if printMsgs {
			fmt.Println("Message for queue:", q, "Next Receive: ", inmsg)
		}

		firstSub = input.Message.Headers.Value("subscription")
		if first {
			if firstSub == "" {
				panic("first subscription header is empty")
			}
			fmt.Println("queue:", q, "FirstSub1: ", firstSub)
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
			fmt.Println("queue:", q, "FirstSub2:", firstSub, "goteof")
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
	uh := stompngo.Headers{"id", firstSub,
		"destination", q}
	error = c.Unsubscribe(uh)
	if error != nil {
		log.Fatalf("unsub error: %v\n", error)
	}
	wg.Done()
	fmt.Println("Receives complete for:", q)
}

func main() {
	fmt.Println("Receiver Start...")

	// create a net.Conn, and pass that into Connect
	p := os.Getenv("STOMP_PORT")
	if p == "" {
		p = "61613"
	}
	nc, error := net.Dial("tcp", hap+p)
	if error != nil {
		log.Fatal(error)
	}
	// Connect
	ch := stompngo.Headers{"login", "getter", "passcode", "recv1234"}
	ch = ch.Add("accept-version", "1.0") // 1.0 only
	ch = ch.Add("host", host)
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
	fmt.Println("Receiver done with wait")
	// Disconnect
	nh := stompngo.Headers{}
	error = c.Disconnect(nh)
	if error != nil {
		log.Fatalf("Receiver discerr %v\n", error)
	}
	// Sanity check for spurious errors
	select {
	case v := <-c.MessageData:
		fmt.Printf("Receiver frame2: %s\n", v.Message.Command)
		fmt.Printf("Receiver header2: %v\n", v.Message.Headers)
		fmt.Printf("Receiver data2: %s\n", string(v.Message.Body))
	default:
		fmt.Println("Receiver Nothing to show")
	}
	// Network close
	nc.Close()
	fmt.Println("Receiver done nc.Close()")
	ngor := runtime.NumGoroutine()
	fmt.Printf("Receiver ngor: %v\n", ngor)
	fmt.Println("Receiver End... numq:", mq, "Num received:", numRecv)
}
