/*
Demo generating a CPU profile.
*/
package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func runner() {
	for i := 0; i < 1000000000; i++ {
		for j := 0; j < 30; j++ { //
			_ = i + j
		}
	}
}

func main() {
	s := time.Now()
	fmt.Println("Start", s)
	//
	f, err := os.Create("./cpu.prof")
	if err != nil {
		panic(err)
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic(err)
	}
	runner()
	pprof.StopCPUProfile()
	//
	err = f.Close()
	if err != nil {
		panic(err)
	}
	e := time.Now()
	fmt.Println("End", s)
	fmt.Println("Elapsed", e.Sub(s))
}
