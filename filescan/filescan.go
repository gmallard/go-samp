/*
A demonstration of using a bufio scanner.  Requires go 1.1.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hi")
	f, e := os.Open("./t.txt")
	if e != nil {
		panic(e)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		t := s.Text()
		if t == "" {
			continue
		}
		fmt.Println(t)
	}
	if e = s.Err(); e != nil {
		panic(e)
	}
	fmt.Println("bye")
}
