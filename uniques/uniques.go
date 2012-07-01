/*
Generate long sequences of unique values.
*/
package main

import (
	"crypto/sha1"
	"fmt" //
	"math/rand"
	"net"
	"time"
)

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
	// fmt.Println("Start...")
	a, e := getAllMacs()
	if e != nil {
		panic("allmacs error")
	}
	// fmt.Printf("%s\n", a)

	// Test with:
	// ./uniques | sort | uniq | wc -l

	maxi := 1000000 // one million
	maxi *= 100     // one hundread million, this will run for a while ......

	for i := 1; i <= maxi; i++ {
		n := fmt.Sprintf("%d", time.Now().UnixNano())
		r := fmt.Sprintf("%d", rand.Int63())
		s := sha1.New()
		t := a + "|" + n + "|" + r
		s.Write([]byte(t))
		v := fmt.Sprintf("%x", s.Sum(nil))
		//
		fmt.Printf("%s\n", v)
	}
	// fmt.Println("End...")
}
