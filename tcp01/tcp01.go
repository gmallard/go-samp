//
package main

import (
	"fmt"
	"net"
	"bufio"
)

//
// A TCP 'echo' example.
// 'Tested' using 'telnet'
//
func runReads(tcpConn *net.TCPConn) {
	br := bufio.NewReader(tcpConn)
	for {
		buffer, err := br.ReadBytes('\n')	// '\n' is delimiter
		if err != nil {
			fmt.Printf("Error = %v\n", err)
			panic("wtf04")
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
}

func main() {
	fmt.Println("Start .....")
	//
	// Get a TCP Address
	//
	tcpAddress, err := net.ResolveTCPAddr("localhost:45678")
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
	runReads(tcpConn)
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
