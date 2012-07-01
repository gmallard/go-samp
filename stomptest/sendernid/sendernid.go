/*
Send STOMP messages using https://github.com/gmallard/stompngo and a STOMP 
1.1 broker.
*/
package main

import (
	"fmt" //
	"github.com/gmallard/stompngo"
	"log"
	"net"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup
var printMsgs bool = true
var nmsgs = 10
var qname = "/queue/stompngo.srpub"
var mq = 2
var host = "localhost"
var hap = host + ":"

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
	}
	error = c.Send(eh, "***EOF*** "+q)
	if error != nil {
		log.Fatal(error)
	}
	wg.Done()

}

func main() {
	fmt.Println("Start...")

	//
	nc, error := net.Dial("tcp", hap+os.Getenv("STOMP_PORT"))
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
	fmt.Println("done with wait")
	// Disconnect
	nh := stompngo.Headers{}
	error = c.Disconnect(nh)
	if error != nil {
		log.Fatalf("discerr %v\n", error)
	}
	fmt.Println("done disconnect, start nc.Close()")
	nc.Close()
	fmt.Println("done nc.Close()")
	ngor := runtime.NumGoroutine()
	fmt.Printf("egor: %v\n", ngor)
	select {
	case v := <-c.MessageData:
		fmt.Printf("frame2: %v\n", v)
	default:
		fmt.Println("Nothing to show")
	}
	fmt.Println("End... ngor:", mq, " nmsgs:", nmsgs)
}
