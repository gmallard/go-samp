/*
Go tour, number 45.
*/
package main

// go tour #45: http://tour.golang.org/#45

// You need to have the tour installed locally to compile this example.  
//
// Install by e.g.:
//
// unset GOPATH # YMMV
// go get code.google.com/p/go-tour/gotour
//

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		m[v] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
