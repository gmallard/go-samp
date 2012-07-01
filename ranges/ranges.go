/*
Using range with maps and strings.
*/
package main

import "fmt"

func main() {
	fmt.Println("Start....")
	//
	m := map[string]float32{"1": 1.0, "pi": 3.1415}
	// Key, Value
	for key, value := range m {
		fmt.Printf("key %s, value %g\n", key, value)
	}
	// Key only
	for key := range m {
		fmt.Printf("key %s\n", key)
	}
	// Characters.
	s := "[\u00ff\u754c]"
	for i, c := range s {
		fmt.Printf("%d:%c ", i, c)
	}
	//
	t := "\u754c"
	fmt.Printf("\nLen: %d\n", len(t))
	fmt.Println("End....")
}
