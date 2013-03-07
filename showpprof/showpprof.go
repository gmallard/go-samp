/*
Display available profiles.
*/
package main

import (
	"log"
	"os"
	"runtime/pprof"
)

var logger = log.New(os.Stdout, "SHOWPPROF ", log.Ldate|log.Lmicroseconds|log.Lshortfile)

func main() {
	logger.Println("Start...")
	//
	profs := pprof.Profiles()
	for _, p := range profs {
		logger.Println(p.Name())
	}
	logger.Println("End...")
}
