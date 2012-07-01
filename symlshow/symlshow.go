/*
Test if a path is a symbolic link.
*/
package main

// Show if a given path is actually a symbolic link

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println(os.Args[1])
	p, _ := filepath.EvalSymlinks(os.Args[1])
	fmt.Println(p)
	fmt.Printf("SYMBOLIC=%v\n", os.Args[1] != p)
}
