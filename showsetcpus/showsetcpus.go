/*
Show retrieving NumCPU and setting GOMAXPROCS.
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	nc := runtime.NumCPU()
	fmt.Println("NumCPUs present:", nc)
	nrc := runtime.GOMAXPROCS(-1)
	fmt.Println("GOMAXPROCS Default:", nrc)
	_ = runtime.GOMAXPROCS(nc)   // Set
	nrc = runtime.GOMAXPROCS(-1) // Retrieve again
	fmt.Println("GOMAXPROCS After Final Set:", nrc)
}
