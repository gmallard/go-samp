/*
   Demo stack trace(s).package stacktrace
*/
package main

import (
	"fmt"
	"runtime/debug"
	"strings"
)

func showStackTrace() {
	// Not very pretty, but ..... pretty useful at times.
	debug.PrintStack()
}

func showStackTrace2(sup int) {
	// Sligtly more pretty, but most of the time just as useful.
	// This could be embellished of course.
	s := string(debug.Stack())
	ls := strings.SplitN(s, "\n", -1)
	for i := sup; i < len(ls); i++ {
		fmt.Printf("%s\n", strings.Trim(ls[i], " \t"))
	}
}

func sublevelb() {
	fmt.Println("sublevelb starts")
	fmt.Println("===============================================")
	showStackTrace() // One way
	fmt.Println("===============================================")
	// 5 seems like a good suppress number (from minimal experimentation)
	showStackTrace2(5) // Or another
	fmt.Println("sublevelb ends")
}

func sublevela() {
	fmt.Println("sublevela starts")
	sublevelb()
	fmt.Println("sublevela ends")
}

func main() {
	fmt.Println("hi")
	sublevela()
	fmt.Println("bye")
}
