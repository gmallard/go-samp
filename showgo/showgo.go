/*
Display go information.
*/
package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("Start....")
	fmt.Printf("GOROOT: %s\n", runtime.GOROOT())
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Version: %s\n", runtime.Version())
	//
	fmt.Println()
	fmt.Printf("TempDir: %s\n", os.TempDir())
	fmt.Printf("PathSep: %s\n", string(os.PathSeparator))
	fmt.Printf("PathListSep: %s\n", string(os.PathListSeparator))
	fmt.Printf("DevNull: %s\n", os.DevNull)
	fmt.Println("End....")
}
