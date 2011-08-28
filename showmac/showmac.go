// Show MAC interfaces on this system

package main

import (
	"fmt" //
	"net"
)

func main() {
	// 
	ifaces, err := net.Interfaces()
	if err != nil {
		// fmt.Printf("Error: %v\n", err)
		panic("ooops 1")
	}
	// fmt.Printf("Interfaces: %v\n", ifaces)
	for _, face := range ifaces {
		fmt.Println("Interface")
		fmt.Println("---------")
		fmt.Printf("Name: %s\n", face.Name)
		fmt.Printf("MAC: %s\n", face.HardwareAddr)
		fmt.Println("")
	}
}
