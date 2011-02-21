//
package main

import (
	"fmt"
	"net"
)

//
// 'Tested' using 'telnet'
//
func runReads(tcpConn *net.TCPConn) {
	for {
		var buffer = make([]byte, 256)
		bytesRead, err := tcpConn.Read(buffer)
		if err != nil {
			fmt.Printf("Error = %v\n", err)
			panic("wtf04")
		}
		//
		fmt.Println("Bytes Read", bytesRead)
		var data = string(buffer[0:bytesRead])
		fmt.Printf("Data Read: |%q|\n", data)
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
