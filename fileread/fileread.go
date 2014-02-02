/*
A text file reader.
*/
package main

//
// A text file reader.
//
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Start....")

	fname := "./data.txt"
	if os.Getenv("INFILE") != "" {
		fname = os.Getenv("INFILE")
	}
	// Open
	// f is *File.
	f, err := os.OpenFile(fname, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("\nOpen Error => %s\n\n", err)
		os.Exit(1)
	}

	// Read lines
	reader := bufio.NewReader(f) // Buffered reader
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("Line Detail: |%q|\n", line)
		fmt.Printf("Line as a string: |%s|\n", line)
	}

	// Close
	err = f.Close()
	if err != nil {
		fmt.Printf("\nClose Error => %s\n\n", err)
		os.Exit(1)
	}
	//
	fmt.Println("End....")
}
