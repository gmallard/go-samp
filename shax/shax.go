// Create a sha1 sum of a string.

package main

import (
	"fmt" //
	"crypto/sha1"
	"crypto/sha256"
)

func main() {
	fmt.Println("Start...")
	s := "Some string A."
	sha1 := sha1.New()
	sha1.Write([]byte(s))
	ss := fmt.Sprintf("%x", sha1.Sum())
	fmt.Printf("%s\n", ss)
	sha256 := sha256.New()
	sha256.Write([]byte(s))
	ss = fmt.Sprintf("%x", sha256.Sum())
	fmt.Printf("%s\n", ss)
	fmt.Println("End...")
}
