// gostompgo demo

package main

import (
	"fmt" //
	"net"

	//	"rand"
	"stomp"
	"strconv"
	"strings"
	//	"time"
	"sync"
	//	"time"
)

var wgsend sync.WaitGroup
var wgrecv sync.WaitGroup
var wgboth sync.WaitGroup
var printMsgs bool = true
// var handp string = "localhost:61613"	// 1.0 server
var handp string = "localhost:62613" // 1.1 server
var nmsgs = 10
var qname = "/queue/gostomp.sendrcv.seq"
var mq = 2 //

var msg_build_ms int64 = 50 // max ms to build a message
var msg_proc_ms int64 = 259 // max ms to process a message


func getNanoSecondsFromMillis(m int64) (n int64) {
	m = n * 1000000 // ms -> ns
	return m
}

const (
	stagger = 1e9 / 4 // Consume some time building and processing messages
)

func recMessages(c *stomp.Connection, q string, k int) {

	var error error
	ks := fmt.Sprintf("%d", k)

	// Receive phase
	headers := stomp.Headers{"destination", q}

	//	sh := headers.Clone()
	sh := headers.Add("id", q)

	fmt.Println("start subscribe", q)
	sc, error := c.Subscribe(sh)
	fmt.Println("end subscribe", q)
	if error != nil {
		panic(error)
	}
	//	if sc != nil {
	//		panic("sc is not nil")
	//	}
	//	for input := range c.MessageData {
	for input := range sc {
		inmsg := input.Message.BodyString()
		if printMsgs {
			fmt.Println("Receive:", q, " / ", inmsg)
		}
		if inmsg == "***EOF***" {
			break
		}
		//
		// time.Sleep(rand.Int63n(stagger) / 2)
		if !strings.HasPrefix(inmsg, ks) {
			fmt.Printf("bad prefix: [%v], [%v], [%v]\n", q, inmsg, ks)
			panic("bad prefix ....")
		}

	}
	fmt.Println("quit for", q)
	error = c.Unsubscribe(headers)
	fmt.Println("end unsubscribe", q)
	if error != nil {
		panic(error)
	}
	wgrecv.Done()
}

func sendMessages(c *stomp.Connection, q string, n int, k int) {

	var error error
	ks := fmt.Sprintf("%d", k)

	// Send
	eh := stomp.Headers{"destination", q} // Extra headers
	for i := 1; i <= n; i++ {
		m := ks + " gostomp message #" + strconv.Itoa(i)
		if printMsgs {
			fmt.Println("Send:", q, " / ", m)
		}
		error = c.Send(eh, m)
		if error != nil {
			panic(error)
		}
		//
		// time.Sleep(rand.Int63n(stagger))
	}

	error = c.Send(eh, "***EOF***")
	if error != nil {
		// Handle error properly
	}
	wgsend.Done()

}

// Test multiple go routines - SEND
func BenchmarkMultipleGoRoutinesSend() {

	// SEND Phase

	// create a net.Conn, and pass that into Connect
	nc, error := net.Dial("tcp", handp)
	if error != nil {
		panic(error)
	}

	// Connect
	ch := stomp.Headers{"login", "guest", "passcode", "guest"}
	c, error := stomp.Connect(nc, ch)
	if error != nil {
		panic(error)
	}

	for i := 1; i <= mq; i++ {
		qn := fmt.Sprintf("%d", i)
		wgsend.Add(1)
		go sendMessages(c, qname+qn, nmsgs, i)

	}
	wgsend.Wait()

	// Disconnect
	nh := stomp.Headers{}
	error = c.Disconnect(nh)
	if error != nil {
		panic(error)
	}

	fmt.Println("Done with SENDs ....")
	wgboth.Done()
}

// Test multiple go routines - RECEIVE
func BenchmarkMultipleGoRoutinesRecv() {

	// RECEIVE Phase

	// create a net.Conn, and pass that into Connect
	nc2, error := net.Dial("tcp", handp)
	if error != nil {
		// Handle error properly
		panic(error)
	}

	// Connect
	ch := stomp.Headers{"login", "guest", "passcode", "guest"}
	c, error := stomp.Connect(nc2, ch)
	if error != nil {
		panic(error)
	}

	for i := 1; i <= mq; i++ {
		qn := fmt.Sprintf("%d", i)
		// fmt.Println("adding", qname + qn)
		wgrecv.Add(1)
		go recMessages(c, qname+qn, i)
	}
	wgrecv.Wait()
	// Disconnect
	nh := stomp.Headers{}
	error = c.Disconnect(nh)
	if error != nil {
		panic(error)
	}
	wgboth.Done()
}

func main() {

	wgboth.Add(2)
	go BenchmarkMultipleGoRoutinesRecv()
	go BenchmarkMultipleGoRoutinesSend()
	wgboth.Wait()
	println("main done")

}
