// Create a sha1 sum of a string.

package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt" //
)

func main() {
	fmt.Println("Start...")
	s := "Some string A."
	sha1 := sha1.New()
	sha1.Write([]byte(s))
	ss := fmt.Sprintf("%x", sha1.Sum(nil))
	fmt.Printf("%s\n", ss)
	sha256 := sha256.New()
	sha256.Write([]byte(s))
	ss = fmt.Sprintf("%x", sha256.Sum(nil))
	fmt.Printf("%s\n", ss)
	fmt.Println("End...")
}
