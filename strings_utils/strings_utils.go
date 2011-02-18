//
package main

import (
	"fmt"
	"strings"
)

func showSplits() {
	astr := "abcd|def|ghij"
	fmt.Printf("Data: %s\n", astr)
	sseps := strings.Split(astr, "|", 2)
	fmt.Printf("Seps Len: %d\n", len(sseps))
	for _, sep_line := range sseps {
		fmt.Printf("SepLine: $%s$\n", sep_line)
	}
	fmt.Println()
	//
	astr = "abcd|def|ghij|"
	fmt.Printf("Data: $%s$\n", astr)
	sseps = strings.Split(astr, "|", 2)
	fmt.Printf("Seps Len: %d\n", len(sseps))
	for _, sep_line := range sseps {
		fmt.Printf("SepLine: $%s$\n", sep_line)
	}
	fmt.Println()
	//
	astr = "abcd|def|ghij"
	fmt.Printf("Data: $%s$\n", astr)
	sseps = strings.SplitAfter(astr, "|", 2)
	fmt.Printf("Seps Len: %d\n", len(sseps))
	for _, sep_line := range sseps {
		fmt.Printf("SepLine: $%s$\n", sep_line)
	}
	fmt.Println()
	//
	astr = "abcd|def|ghij|"
	fmt.Printf("Data: $%s$\n", astr)
	sseps = strings.SplitAfter(astr, "|", 2)
	fmt.Printf("Seps Len: %d\n", len(sseps))
	for _, sep_line := range sseps {
		fmt.Printf("SepLine: $%s$\n", sep_line)
	}
	fmt.Println()
	//
	astr = "abcd\ndef\nghij\n\ndata"
	fmt.Printf("Data: $%s$\n", astr)
	sseps = strings.Split(astr, "\n\n", 2)
	fmt.Printf("Seps Len: %d\n", len(sseps))
	for _, sep_line := range sseps {
		fmt.Printf("SepLine: $%s$\n", sep_line)
	}
	fmt.Println()
}

func main() {
	fmt.Println("Start....")
	//
	showSplits()
	//
	fmt.Println("End....")
}
