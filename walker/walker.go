/*
   Demo filepath Walk
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	fc = 0
	dc = 0
)

func main() {
	fmt.Println("hi")
	hb := os.Getenv("HOME")
	fmt.Println("HOME", hb)
	err := filepath.Walk(hb, myWalker)
	if err != nil {
		fmt.Println("MAIN ERROR:", err)
	}
	fmt.Println("Dir Count", dc)
	fmt.Println("File Count", fc)
	fmt.Println("bye")
}

func myWalker(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("%s %v\n", "Walker Error: ", err)
		return nil
	}
	if info.IsDir() {
		fmt.Println("Directory:", info.Name())
		dc++
	} else {
		fmt.Println("File:", info.Name())
		fc++
	}
	return nil
}
