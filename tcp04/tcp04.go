/*
How not to use go routines for reading.
*/
package main

import (
	"fmt"
	"net"
	"strings"
)

// Demo goroutine start, but main goroutine exit, which kills the connection.
// Not useful, but indicates how connections work.
// 'Tested' using 'telnet localhost 45678'
func getData(tcpConn *net.TCPConn, done chan bool) {
	for {
		var buffer = make([]byte, 256)
		fmt.Println("Starting Read ...")
		bytesRead, err := tcpConn.Read(buffer)
		if err != nil {
			fmt.Printf("Error = %v\n", err)
			break
		}
		//
		fmt.Println("Bytes Read", tcpConn, bytesRead)
		bufData := buffer[0:bytesRead]
		fmt.Println("Buffer", tcpConn, bufData)
		var data = string(bufData)
		fmt.Println("Data Read", tcpConn, data)
		//
		if strings.HasPrefix(data, "quit") {
			break
		}
	}
	//
	fmt.Println("Starting Close() 1")
	err := tcpConn.Close()
	if err != nil {
		fmt.Printf("Error = %v\n", err)
		panic("wtf05")
	}
	//
	done <- true
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
	// Accept connections.
	//
	tcpConn, err := listener.AcceptTCP()
	if err != nil {
		fmt.Printf("Error = %v\n", err)
		panic("wtf03")
	}
	fmt.Printf("connection = %v\n", tcpConn)
	waitFor := make(chan bool)
	go getData(tcpConn, waitFor)
	// Do not wait ......
	// <- waitFor
	//
	fmt.Println("Starting Close() 2")
	err = listener.Close()
	if err != nil {
		fmt.Printf("Error = %v\n", err)
		panic("wtf06")
	}
	//
	fmt.Println("End .....")
}
