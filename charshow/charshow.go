/*
Playing with UTF8 strings.
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func hrune(r rune) string {
	return fmt.Sprintf("%08x", r)
}

func main() {
	fmt.Println("hi")
	s := "𐒁̸"
	fmt.Println(s, len(s), utf8.RuneCountInString(s))
	s = "汉"
	fmt.Println(s, len(s), utf8.RuneCountInString(s))
	s = "𐒁̸汉"
	fmt.Println(s, len(s), utf8.RuneCountInString(s))
	//
	for i, r := range s {
		fmt.Println(i, hrune(r), " ", string(r), utf8.RuneLen(r))
	}
}
