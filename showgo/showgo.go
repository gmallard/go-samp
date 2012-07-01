/*
Display go information.
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Start....")
	fmt.Printf("GOROOT: %s\n", runtime.GOROOT())
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Version: %s\n", runtime.Version())
	fmt.Println("End....")
}
