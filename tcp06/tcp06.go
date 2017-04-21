/*
TCP demonstration, multiple clients using go routines.
*/
package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

var shutdown = false

// Demo multiple echo clients.
// 'Test' using 'telnet localhost 45678' from multiple open terminals.
func getData(c *net.TCPConn) {
	var buffer = make([]byte, 4096)
	a := c.RemoteAddr()
	fmt.Printf("Net: %s, Remote: %s\n", a.Network(), a.String())
	for {
		if shutdown {
			fmt.Println(c, "Shutdown")
			break
		}
		fmt.Println(c, "Starting Read ...")
		_ = c.SetReadDeadline(time.Now().Add(1 * time.Minute))
		bytesRead, err := c.Read(buffer)
		if err != nil && err.(net.Error).Timeout() {
			fmt.Printf("Timeout Read: %v\n", err)
			continue
		}
		if err != nil {
			fmt.Printf("Error = %v\n", err)
			break
		}
		//
		fmt.Println(c, "Bytes Read", bytesRead)
		bufData := buffer[0:bytesRead]
		fmt.Println(c, "Raw Buffer", bufData)
		var data = string(bufData)
		fmt.Println(c, "Data Read", data)
		//
		s := "Echo: " + data
		if _, err := c.Write([]byte(s)); err != nil {
			fmt.Printf("Error = %v\n", err)
			panic("wtfwr01")
		}
		//
		if strings.HasPrefix(data, "quit") {
			break
		}
		//
		if strings.HasPrefix(data, "shutdown") {
			shutdown = true
		}
	}
	//
	fmt.Println(c, "Starting Close() 1")
	err := c.Close()
	if err != nil {
		fmt.Printf("Error = %v\n", err)
		panic("wtf05")
	}
	//
}

func main() {
	fmt.Println("Start .....")
	//
	// Get a TCP Address
	//
	a, err := net.ResolveTCPAddr("", "localhost:45678")
	if err != nil {
		fmt.Printf("Error = %v\n", err)
		panic("wtf01")
	}
	fmt.Printf("a = %v\n", a)
	//
	// Get a TCP Listener
	//
	l, err := net.ListenTCP("tcp", a)
	if err != nil {
		fmt.Printf("Error = %v\n", err)
		panic("wtf02")
	}
	fmt.Printf("l = %v\n", l)
	//
	// Accept any number of connections.
	//
	for {
		if shutdown {
			fmt.Println(l, "Listener shutdown")
			break
		}
		_ = l.SetDeadline(time.Now().Add(2 * time.Minute))
		c, err := l.AcceptTCP()
		if err != nil && err.(net.Error).Timeout() {
			fmt.Printf("Timeout Accept: %v\n", err)
			continue
		}
		if err != nil {
			fmt.Printf("Error = %v\n", err)
			panic("wtf03")
		}
		fmt.Printf("New Connection = %v\n", c)
		go getData(c)
	}
	//
	fmt.Println("Starting Close() 2")
	err = l.Close()
	if err != nil {
		fmt.Printf("Error = %v\n", err)
		panic("wtf06")
	}
	//
	fmt.Println("End .....")
}
