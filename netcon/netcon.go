/*
Use a net.Conn connection, which is an interface.
*/
package main

import (
	"fmt"  //
	"net"  //
	"os"   //
	"time" //
)

func main() {
	fmt.Println("Start...")
	conn, err := net.Dial("tcp", "localhost:54321")
	//
	if err != nil {
		fmt.Printf("Dial Error received: %v\n", err)
		os.Exit(4)
	}
	//
	time.Sleep(5 * time.Second) // 5 seconds
	//
	err = conn.Close()
	if err != nil {
		fmt.Printf("Close Error received: %v\n", err)
		os.Exit(4)
	}
	//
	fmt.Println("End...")
}
