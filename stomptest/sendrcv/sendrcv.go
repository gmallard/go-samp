// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt" //
  "net"
  "os"
  "strconv"
	"stomp"
	"sync"
)

// Not really a benchmark.  This will cause a panic in the gostomp code
// as of 09/09/2011.  Has to do with newlines between frames on the wire.
// Many stomp brokers put extraneous newlines between frames, and in Stomp 1.1
// this is actually used for the heartbeat facility.  The panic occurs only
// under some load and only with particular broker implementations.

var wgsend sync.WaitGroup
var wgrecv sync.WaitGroup
var printMsgs bool = true
var handp string = "localhost:61613"

func recMessages(c *stomp.Conn, q string) {

	var error os.Error

	// Receive phase
  headers := stomp.Header{"destination": q}
	_, error = c.Subscribe(headers)
	if error != nil {
		// Handle error properly
	}
	for input := range c.Stompdata {
    inmsg := string(input.Message.Data)
    if printMsgs {
  		fmt.Println("queue:", q, "Next Receive: ", inmsg)
    }
		if inmsg == "***EOF***" {
			break
		}
	}
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
			fmt.Println("queue:", q, "msg:", m)
		}
	  error = c.Send(eh, m)
	  if error != nil {
		  // Handle error properly
	  }
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
		// Handle error properly
	}

  // Connect
	ch := stomp.Header{"login": "userid", "passcode": "abcd1234"}
	c, error := stomp.Connect(nc, ch)
	if error != nil {
		// Handle error properly
	}

  nmsgs := 200
	qname := "/queue/gostomp.srpub"
	mq := 25	// Too many of these triggers the panic during RECEIVE

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
		// Handle error properly
	}

	fmt.Println("Done with SENDs ....")

}

// Test multiple go routines - SEND
func BenchmarkMultipleGoRoutinesRecv() {

	// RECEIVE Phase

  // create a net.Conn, and pass that into Connect
	fmt.Println("DB01")
	nc2, error := net.Dial("tcp", handp)
	if error != nil {
		// Handle error properly
		fmt.Println("DB02", error)
	}
		fmt.Println("DB03")

	fmt.Println("Done with RECEIVE net conn ....")

  // Connect
	ch := stomp.Header{"login": "userid", "passcode": "abcd1234"}
	c, error := stomp.Connect(nc2, ch)
	if error != nil {
		// Handle error properly
	}

	fmt.Println("Done with RECEIVE stomp.Connect ....")

	qname := "/queue/gostomp.srpub"
	mq := 25	// Too many of these triggers the panic during RECEIVE

	for i := 1; i <= mq; i++ {
		qn := fmt.Sprintf("%d", i)
		wgrecv.Add(1)
		go recMessages(c, qname + qn)
	}
	wgrecv.Wait()

  // Disconnect
  nh := stomp.Header{}
	error = c.Disconnect(nh)
	if error != nil {
		// Handle error properly
	}

}

func main() {

	BenchmarkMultipleGoRoutinesSend()
	BenchmarkMultipleGoRoutinesRecv()

}

