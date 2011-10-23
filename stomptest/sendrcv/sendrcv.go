// gostomp demo

package main

import (
	"fmt" //
  "net"
  "os"
	"rand"
	"runtime"
  "strconv"
	"stomp"
	"sync"
	"time"
)

var wgsend sync.WaitGroup
var wgrecv sync.WaitGroup
var wgboth sync.WaitGroup
var printMsgs bool = true
// var handp string = "localhost:61613"	// 1.0 server
var handp string = "localhost:62613"	// 1.1 server
var nmsgs = 500
var	qname = "/queue/gostomp.sendrcv.seq"
var	mq = 50	//

var msg_build_ms int64 = 50 // max ms to build a message
var msg_proc_ms int64 = 259 // max ms to process a message


func getNanoSecondsFromMillis(m int64) (n int64) {
	m = n * 1000000 // ms -> ns
	return m
}

func recMessages(c *stomp.Conn, q string) {

	var error os.Error

	// Receive phase
  headers := stomp.Header{"destination": q, "id": q}
	fmt.Println("start subscribe", q)
	sc, error := c.Subscribe(headers)
	fmt.Println("end subscribe", q)
	if error != nil {
		panic(error)
	}
	k := 0
	for input := range sc {
    inmsg := string(input.Message.Data)
    if printMsgs {
  		fmt.Println("Receive:", q, " / ", inmsg)
    }
		if inmsg == "***EOF***" {
			break
		}
		k++
		if k % 10 == 0 {
			fmt.Println("Receive:", q, "yields")
			runtime.Gosched()	// yield
		}
		time.Sleep(rand.Int63n(getNanoSecondsFromMillis(msg_proc_ms)))
	}
	fmt.Println("quit for", q)
	wgrecv.Done()
}

func sendMessages(c *stomp.Conn, q string, n int, k int) {

	var error os.Error
	ks := fmt.Sprintf("%d", k)

  // Send
  eh := stomp.Header{"destination": q} // Extra headers
  for i := 1; i <= n; i++ {
		m := ks + " gostomp message #" + strconv.Itoa(i)
		if printMsgs {
			fmt.Println("Send:", q, " / ", m)
		}
	  error = c.Send(eh, m)
	  if error != nil {
			panic(error)
	  }
		if i % 10 == 0 {
			fmt.Println("Send:", q, "yields")
			runtime.Gosched()	// yield
		}
		time.Sleep(rand.Int63n(getNanoSecondsFromMillis(msg_build_ms)))
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
	ch := stomp.Header{"login": "guest", "passcode": "guest"}
	c, error := stomp.Connect(nc, ch)
	if error != nil {
		panic(error)
	}

	for i := 1; i <= mq; i++ {
		qn := fmt.Sprintf("%d", i)
		wgsend.Add(1)
		go sendMessages(c, qname + qn, nmsgs, i)

	}
	wgsend.Wait()

  // Disconnect
  nh := stomp.Header{}
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
	ch := stomp.Header{"login": "guest", "passcode": "guest"}
	c, error := stomp.Connect(nc2, ch)
	if error != nil {
		panic(error)
	}

	for i := 1; i <= mq; i++ {
		qn := fmt.Sprintf("%d", i)
		// fmt.Println("adding", qname + qn)
		wgrecv.Add(1)
		go recMessages(c, qname + qn)
	}
	wgrecv.Wait()

  // Disconnect
  nh := stomp.Header{}
	error = c.Disconnect(nh)
	if error != nil {
		panic(error)
	}
	fmt.Println("Done with RECEIVEs ....")
	wgboth.Done()
}

func main() {

	wgboth.Add(2)
	go BenchmarkMultipleGoRoutinesRecv()
	go BenchmarkMultipleGoRoutinesSend()
	wgboth.Wait()

}

