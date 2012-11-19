/*
Demonstrate UTF-8 oddities with characters "σ" and "ς".
*/
package main

import "fmt"
import "strings"

func main() {
	// The Greek Alphabet in Upper Case
	g := "ΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡΣΤΥΦΧΨΩ"
	fmt.Println(g)
	fmt.Println()
	fmt.Println(strings.ToLower(g))
	fmt.Println()

	// But what about the ς versus σ thing?
	v := "ς"
	fmt.Println("c-cedilla", v)
	fmt.Println("c-cedilla UPPER", strings.ToUpper(v))
	fmt.Println("c-cedilla flip", strings.ToLower(strings.ToUpper(v)))
	//
	v = "σ"
	fmt.Println("sigma", v)
	fmt.Println("sigma UPPER", strings.ToUpper(v))
	fmt.Println("sigma flip", strings.ToLower(strings.ToUpper(v)))

	fmt.Println()

	// But what about the Ç versus σ thing?
	// Is Ç really like ς ????
	v = "Ç"
	fmt.Println("c-cedillau", v)
	fmt.Println("c-cedillau UPPER", strings.ToUpper(v))
	fmt.Println("c-cedillau flip", strings.ToLower(strings.ToUpper(v)))
	//
	v = "σ"
	fmt.Println("sigmau", v)
	fmt.Println("sigmau UPPER", strings.ToUpper(v))
	fmt.Println("sigmau flip", strings.ToLower(strings.ToUpper(v)))

}
