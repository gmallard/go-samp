/*
	Dump a file in hex.
*/
package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("DumpFile Starts....")

	fname := os.Args[1]
	if os.Getenv("INFILE") != "" {
		fname = os.Getenv("INFILE")
	}
	fmt.Printf("Dumpfile Input: %s\n", fname)

	// Open
	f, err := os.OpenFile(fname, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("\nDumpFile Open Error => %s\n\n", err)
		os.Exit(1)
	}

	// Dump
	buff, err := ioutil.ReadAll(f)
	fmt.Printf("%s", hex.Dump(buff))

	// Close
	err = f.Close()
	if err != nil {
		fmt.Printf("\nClose Error => %s\n\n", err)
		os.Exit(1)
	}
	//
	fmt.Println("DumpFile Ends....")
}
