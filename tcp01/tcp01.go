/*
A simple TCP echo server.
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

//
// A TCP 'echo' example.
// Demonstrates using a timeout, and a 'graceful' shutdown if one occurs.
// 'Tested' using 'telnet localhost 45678'
//
func runReads(tcpConn *net.TCPConn) bool {
	br := bufio.NewReader(tcpConn)
	for {
		// Set a timeout value, which needs to be set before each and every read.
		d := time.Duration(30 * 1e9) // 30 seconds
		w := time.Now()              // from now
		w = w.Add(d)
		tcpConn.SetReadDeadline(w) // Set the deadline
		//
		buffer, err := br.ReadBytes('\n') // '\n' is delimiter
		// If the read times out, this prints something like:
		// Error = read tcp 127.0.0.1:57609: resource temporarily unavailable
		if err != nil {
			fmt.Printf("Error = %v\n", err)
			return false
			// panic("wtf04")
		}
		//
		fmt.Printf("Bytes Read: %d\n", len(buffer))
		var data = string(buffer)
		fmt.Printf("Data Read: |%q|\n", data)

		// This is now an 'echo' example.
		out := "echo: " + data
		tcpConn.Write([]byte(out))

		// The \r in this data from telnet is a bit surprising ...
		if data == "quit\r\n" {
			fmt.Println("Breaking....")
			break
		}
	}
	return true
}

func main() {
	fmt.Println("Start .....")
	//
	// Get a TCP Address
	//
	tcpAddress, err := net.ResolveTCPAddr("", "localhost:45678")
	if err != nil {
		fmt.Printf("Error = %v\n", err)
		panic("wtf01")
	}
	fmt.Printf("tcpAddress = %v\n", tcpAddress)
	//
	// Get a TCP Listener
	//
	listener, err := net.ListenTCP("tcp", tcpAddress)
	if err != nil {
		fmt.Printf("Error = %v\n", err)
		panic("wtf02")
	}
	fmt.Printf("listener = %v\n", listener)
	//
	// Accept a connection.
	//
	tcpConn, err := listener.AcceptTCP()
	if err != nil {
		fmt.Printf("Error = %v\n", err)
		panic("wtf03")
	}
	fmt.Printf("connection = %v\n", tcpConn)
	//
	checkVal := runReads(tcpConn)
	fmt.Println("CheckVal", checkVal)
	//
	err = tcpConn.Close()
	if err != nil {
		fmt.Printf("Error = %v\n", err)
		panic("wtf05")
	}
	//
	err = listener.Close()
	if err != nil {
		fmt.Printf("Error = %v\n", err)
		panic("wtf06")
	}
	//
	fmt.Println("End .....")
}
