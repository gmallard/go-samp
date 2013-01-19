/*
An example of using recover, from http://play.golang.org/p/m96skGjRjo
*/
package main

import (
	"log"
	"runtime/debug"
)

func deep2(n int) {
	defer func() {
		log.Printf("Resource %d closed", n)
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	panic("I'm afraid I can't do that")
}

func deep1(n int) {
	defer func() {
		log.Printf("Resource %d closed", n)
		if err := recover(); err != nil {
			panic(err)
		}
	}()
	deep2(n + 1)
}

func main() {
	defer func() {
		log.Print("Main exiting")
		if err := recover(); err != nil {
			log.Fatalf("Stack trace:\n%s----\n%s", debug.Stack(), err)
		}
	}()

	log.Println("Main started")
	deep1(1)
}

