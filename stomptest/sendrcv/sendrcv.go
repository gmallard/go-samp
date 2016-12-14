/*
Send and receive STOMP messages using https://github.com/gmallard/stompngo using multiple
go routines.
*/
package main

import (
	"fmt" //
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gmallard/stompngo"
)

var wgsend sync.WaitGroup
var wgrecv sync.WaitGroup
var wgboth sync.WaitGroup
var printMsgs bool = true

var host = "localhost"
var hap = host + ":" //
var nmsgs = 100
var qname = "/queue/stompngo.sendrcv.seq"
var nq = 10 //

var seed int64       // The seed value (time.Now().UnixNano())
var src rand.Source  // The source
var rgptr *rand.Rand // Pointer to the generator

func getStagger(minns, maxns int64) int64 {
	return minns + rgptr.Int63n(maxns-minns)
}

func recMessages(c *stompngo.Connection, q string, k int) {

	var error error
	ks := fmt.Sprintf("%d", k)

	// Receive phase
	headers := stompngo.Headers{"destination", q}
	sh := headers.Add("id", q)
	//
	log.Println("start subscribe", q)
	sc, error := c.Subscribe(sh)
	log.Println("end subscribe", q)
	if error != nil {
		log.Fatal(error)
	}
	for input := range sc {
		inmsg := input.Message.BodyString()
		if printMsgs {
			log.Println("Receive:", q, " / ", inmsg)
		}
		if inmsg == "***EOF***" {
			break
		}
		if !strings.HasPrefix(inmsg, ks) {
			log.Printf("bad prefix: [%v], [%v], [%v]\n", q, inmsg, ks)
			log.Fatal("bad prefix ....")
		}
		//
		d := time.Duration(getStagger(1e9/10, 1e9/5))
		time.Sleep(d)
	}
	log.Println("quit for", q)
	error = c.Unsubscribe(headers)
	log.Println("end unsubscribe", q)
	if error != nil {
		log.Fatal(error)
	}
	wgrecv.Done()
}

func sendMessages(c *stompngo.Connection, q string, n int, k int) {

	var error error
	ks := fmt.Sprintf("%d", k)
	// Send
	eh := stompngo.Headers{"destination", q} // Extra headers
	for i := 1; i <= n; i++ {
		m := ks + " gostomp message #" + strconv.Itoa(i)
		if printMsgs {
			log.Println("Send:", q, " / ", m)
		}
		error = c.Send(eh, m)
		if error != nil {
			log.Fatal(error)
		}
		//
		d := time.Duration(getStagger(1e9/20, 1e9/10))
		time.Sleep(d)
	}
	error = c.Send(eh, "***EOF***")
	if error != nil {
		log.Fatal(error)
	}
	wgsend.Done()
}

// Test multiple go routines - SEND
func BenchmarkMultipleGoRoutinesSend() {

	// SEND Phase
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
	ch := stompngo.Headers{"login", "guest", "passcode", "guest"}
	c, error := stompngo.Connect(nc, ch)
	if error != nil {
		log.Fatal(error)
	}
	for i := 1; i <= nq; i++ {
		qn := fmt.Sprintf("%d", i)
		wgsend.Add(1)
		// All sending go routines share the same connection in this example
		go sendMessages(c, qname+qn, nmsgs, i)

	}
	wgsend.Wait()
	// Disconnect
	nh := stompngo.Headers{}
	error = c.Disconnect(nh)
	if error != nil {
		log.Fatal(error)
	}
	log.Println("Done with SENDs ....")
	wgboth.Done()
}

// Test multiple go routines - RECEIVE
func BenchmarkMultipleGoRoutinesRecv() {

	// RECEIVE Phase
	p := os.Getenv("STOMP_PORT")
	if p == "" {
		p = "61613"
	}
	nc2, error := net.Dial("tcp", hap+p)
	if error != nil {
		// Handle error properly
		log.Fatal(error)
	}
	// Connect
	ch := stompngo.Headers{"login", "guest", "passcode", "guest"}
	c, error := stompngo.Connect(nc2, ch)
	if error != nil {
		log.Fatal(error)
	}
	for i := 1; i <= nq; i++ {
		qn := fmt.Sprintf("%d", i)
		wgrecv.Add(1)
		// All receiving go routines share the same connection in this example
		go recMessages(c, qname+qn, i)
	}
	wgrecv.Wait()
	// Disconnect
	nh := stompngo.Headers{}
	error = c.Disconnect(nh)
	if error != nil {
		log.Fatal(error)
	}
	wgboth.Done()
}

func main() {

	seed = time.Now().UnixNano()
	src = rand.NewSource(seed)
	rgptr = rand.New(src)

	wgboth.Add(2)
	go BenchmarkMultipleGoRoutinesRecv()
	go BenchmarkMultipleGoRoutinesSend()
	wgboth.Wait()
	println("main done")

}
