/*
Short demonstration of signals.
*/
package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

var wg sync.WaitGroup

func handleSignal(c chan os.Signal) {
	s := <-c
	fmt.Println("signal handled:", s)
	wg.Done()
}

func main() {
	fmt.Println("Start...")
	wg.Add(1)

	i := make(chan os.Signal, 1)
	go handleSignal(i)
	signal.Notify(i) // All handleable signals

	wg.Wait()
	fmt.Println("End...")
}
