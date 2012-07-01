/*
Generating sha1 and sha256 sums.
*/
package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt" //
)

func main() {
	fmt.Println("Start...")
	// String from:
	// http://en.wikipedia.org/wiki/WebSocket
	s := "x3JJHMbDL1EzLkh9GBhXDw==258EAFA5-E914-47DA-95CA-C5AB0DC85B11"

	// sha1 stuff ...
	sha1 := sha1.New()
	sha1.Write([]byte(s))
	ss := fmt.Sprintf("%x", sha1.Sum(nil))
	fmt.Printf("%s\n", ss)
	w := "1d29ab734b0c9585240069a6e4e3e91b61da1969"
	fmt.Printf("%s\n", w)
	if ss != w {
		panic("Uh oh, something is not right")
	}
	// ---------------------------------------------------------------------------
	// The base64 encoding part of that post is left as an exercise for now.
	// From the article the base64 result should be:
	// HSmrc0sMlYUkAGmm5OPpG2HaGWk=
	// *not tested*
	// ---------------------------------------------------------------------------
	// sha256 stuff ...
	sha256 := sha256.New()
	sha256.Write([]byte(s))
	ss = fmt.Sprintf("%x", sha256.Sum(nil))
	fmt.Printf("%s\n", ss)
	fmt.Println("End...")
}
