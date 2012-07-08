/*
An example taken from: https://groups.google.com/forum/?fromgroups#!topic/golang-nuts/VdmoZ59jjoE
*/
package main

import "fmt"

//
// See:
// https://groups.google.com/forum/?fromgroups#!topic/golang-nuts/VdmoZ59jjoE
// http://play.golang.org/p/ZJamJZgeEK
//
// A channel shunt (or shovel).
//
func shunt(input <-chan int, output chan<- int) {
	var (
		i   int
		ok  bool
		in  = input
		out chan<- int
	)
	for {
		select {
		case i, ok = <-in:
			if !ok {
				close(output)
				fmt.Println("done")
				return
			}
			fmt.Println("shunt in", i)
			in = nil
			out = output
		case out <- i:
			fmt.Println("shunt out", i)
			in = input
			out = nil
		}
	}

}

func main() {
	input := make(chan int, 100)
	go func() { // Simulate a sender to input in another part of the program
		for i := 0; i < 1000; i++ {
			input <- i
		}
		close(input)
	}()

	acc := make(chan int, 10) // accumulator channel
	go shunt(input, acc)
	for i := range acc {
		fmt.Println("acc read:", i)
	}
	fmt.Println("exit")
}
