package main

import "fmt"

func main() {
	fmt.Println("Start....")
	//
	astring := "A\n\u00ffÇ\u754c"
	fmt.Printf("Len: %d\n", len(astring))
	//
	for char_index, char := range astring {
		var cbytes = []uint8(string(char))
		fmt.Printf("%d:%c: Byte Count=%d\n", char_index, char, len(cbytes))
	}
	//
	fmt.Println()
	fmt.Println("End....")
}
