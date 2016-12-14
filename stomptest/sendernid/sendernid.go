/*
Send STOMP messages using https://github.com/gmallard/stompngo and a STOMP
1.1 broker.
*/
package main

import (
	"fmt" //
	"log"
	"net"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/gmallard/stompngo"
)

var wg sync.WaitGroup
var printMsgs bool = true
var nmsgs = 10
var qname = "/queue/stompngo.srpub"
var mq = 2
var host = "localhost"
var hap = host + ":"

var incrCtl sync.Mutex
var numSend int

func sendMessages(c *stompngo.Connection, q string, n int, k int) {

	var error error

	// Send
	eh := stompngo.Headers{"destination", q} // Extra headers
	for i := 1; i <= n; i++ {
		m := q + " gostomp message #" + strconv.Itoa(i)
		if printMsgs {
			fmt.Println("msg:", m)
		}
		error = c.Send(eh, m)
		if error != nil {
			log.Fatalf("send error: %v\n", error)
		}
		//
		time.Sleep(1e9 / 100) // Simulate message build
		incrCtl.Lock()
		numSend++
		incrCtl.Unlock()
	}
	error = c.Send(eh, "***EOF*** "+q)
	incrCtl.Lock()
	numSend++
	incrCtl.Unlock()
	if error != nil {
		log.Fatal(error)
	}
	wg.Done()
	fmt.Println("Sends complete for:", q)
}

func main() {
	fmt.Println("Sender Start...")

	//
	p := os.Getenv("STOMP_PORT")
	if p == "" {
		p = "61613"
	}
	nc, error := net.Dial("tcp", hap+p)
	if error != nil {
		// Handle error properly
	}
	// Connect
	ch := stompngo.Headers{"login", "putter", "passcode", "send1234",
		"accept-version", "1.1", "host", host}
	c, error := stompngo.Connect(nc, ch)
	if error != nil {
		log.Fatal(error)
	}
	for i := 1; i <= mq; i++ {
		qn := fmt.Sprintf("%d", i)
		wg.Add(1)
		go sendMessages(c, qname+qn, nmsgs, i)

	}
	wg.Wait()
	fmt.Println("Sender done with wait")
	// Disconnect
	nh := stompngo.Headers{}
	error = c.Disconnect(nh)
	if error != nil {
		log.Fatalf("Sender discerr %v\n", error)
	}
	// Sanity check for spurious errors.
	select {
	case v := <-c.MessageData:
		fmt.Printf("frame2: %v\n", v)
	default:
		fmt.Println("Sender Nothing to show")
	}
	fmt.Println("Sender done disconnect, start nc.Close()")
	// Network close
	nc.Close()
	fmt.Println("Sender done nc.Close()")
	ngor := runtime.NumGoroutine()
	fmt.Printf("Sender ngor: %v\n", ngor)
	fmt.Println("Sender End... numq:", mq, "Num sent:", numSend)
}
