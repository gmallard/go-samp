/*
Experiment with string lengths and character sequences.
*/
package main

import "fmt"

func main() {
	fmt.Println("Start....")
	//
	astring := "A\n\u00ffÇ\u754c" // Some UTF-8
	fmt.Printf("AstringLen: %d\n", len(astring))
	//
	b1 := astring[1] // a byte
	fmt.Printf("b1val: %v\n", b1)
	//
	for char_index, char := range astring { // All characters
		cstring := string(char)                      // Convert a character to a string
		fmt.Printf("CstringLen: %d\n", len(cstring)) // len = 1..4 bytes
		var cbytes = []uint8(cstring)                // get the bytes
		// Info
		fmt.Printf("%d:%c: Byte Count=%d\n", char_index, char, len(cbytes))
	}
	//
	fmt.Println()
	fmt.Println("End....")
}
