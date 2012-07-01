/*
Using a mutex for synchronization of go routines.
*/
package main

import (
	"fmt"  //
	"sync" //
	"time" //
)

type gdata struct {
	id            string
	loops, sltime int
}

var runData = []gdata{gdata{"1", 6, 2},
	gdata{"2", 5, 1},
	gdata{"3", 8, 3},
}

var processing, decrCtl sync.Mutex
var numRtns int

func called(cd gdata) {
	for i := 0; i < cd.loops; i++ {
		wnum := i + 1
		fmt.Printf("id %s, waitnum: %d\n", cd.id, wnum)
		time.Sleep(time.Duration(int64(cd.sltime) * 1e9))
	}
	//
	decrCtl.Lock()
	numRtns--
	if numRtns == 0 {
		processing.Unlock()
	}
	decrCtl.Unlock()
	//
	fmt.Println(cd.id, "is done")
}

func main() {
	fmt.Println("Start...")
	//
	numRtns = len(runData)
	processing.Lock()
	//
	for _, curgd := range runData {
		go called(curgd)
	}
	//
	fmt.Println("Starting main wait")
	processing.Lock()
	//
	fmt.Println("End...")
}
