/*
Short demo of base64 encoding/decoding.
*/
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	n := base64.StdEncoding // Use a standard encoding

	// Show some encoded lengths ...
	for i := 1; i <= 16; i++ {
		fmt.Println(i, n.EncodedLen(i))
	}

	fmt.Println()

	// Encode and then decode some strings ...
	s := []string{"A", "AB", "ABC", "ABCD"}
	for _, v := range s {
		b := []byte(v)
		ev := n.EncodeToString(b)
		fmt.Println(v, ev)
		ob, e := n.DecodeString(ev)
		if e != nil {
			panic(e)
		}
		fmt.Println(v, string(ob))
	}
}
