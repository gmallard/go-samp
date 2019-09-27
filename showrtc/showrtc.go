package main

import (
	"log"
	"runtime"
)

/*
Show the constants from the runtime package.
*/
func main() {
	log.Println("runtime.Compiler", runtime.Compiler)
	log.Println("runtime.GOARCH", runtime.GOARCH)
	log.Println("runtime.GOOS", runtime.GOOS)
}
