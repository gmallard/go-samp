/*
Show MAC / NIC information for this system.
*/
package main

import (
	"fmt" //
	"net"
	"strings"
)

func getLocalMac() (result string, error error) {
	result = ""
	error = nil
	//
	ifaces, error := net.Interfaces()
	if error != nil {
		return
	}
	for _, face := range ifaces {
		// Eliminate local and possible Virtual Box interfaces
		if face.Name == "lo" || strings.HasPrefix(face.Name, "vbox") {
			continue
		}
		// Prefer known names for local interfaces
		if strings.HasPrefix(face.Name, "eth") || strings.HasPrefix(face.Name, "en") {
			result = face.HardwareAddr.String()
			break
		}
		// Otherwise, take what we can get
		result = face.HardwareAddr.String()
	}
	//
	return
}

func getAllMacs() (result string, error error) {
	result = ""
	error = nil
	//
	ifaces, error := net.Interfaces()
	if error != nil {
		return
	}
	for _, face := range ifaces {
		//
		result += (face.HardwareAddr.String() + "~")
	}
	//
	return
}

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
	//
	preferred, err := getLocalMac()
	if err != nil {
		panic("ooops 2")
	}
	fmt.Printf("Preferred MAC: %v\n", preferred)
	//
	allmacs, err := getAllMacs()
	if err != nil {
		panic("ooops 3")
	}
	fmt.Printf("All MACs: %v\n", allmacs)
}
