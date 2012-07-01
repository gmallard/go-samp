/*
More symbolic link fun.
*/
package main

import (
	"fmt" //
	"os"  //
)

// Symbolic link gotcha's

func main() {
	fmt.Println("Start...")
	goroot := os.Getenv("GOROOT")
	// In my environment:
	// GOROOT=/home/gmallard/hext/go
	// The 'hext' part is actually a symbolic link -> /ad2/gma/home_ext
	fmt.Printf("GOROOT=%s\n", goroot)
	if err := os.Chdir(goroot); err != nil {
		fmt.Println("chdir error", err)
		os.Exit(1)
	}
	// Note: Getwd documentation says if path is reachable multiple ways,
	// _any_ of the ways could be returned.  Try it.
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("getwd error", err)
		os.Exit(1)
	}
	// And in my particular case, this shows:
	// CWD=/ad2/gma/home_ext/go
	fmt.Printf("CWD=%s\n", cwd)
	fmt.Println("End...")
}
