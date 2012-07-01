/*
Flag package, a short demo.
*/
package main

import (
	"flag"
	"fmt" //
	"time"
)

// Flag variables
var (
	intFlag  int
	strFlag  string
	boolFlag bool
	durFlag  time.Duration
)

// Main initialization, a convenient place to set flags up
func init() {
	flag.IntVar(&intFlag, "fintflag", 1234, "help message for fintflag")
	flag.StringVar(&strFlag, "fstrflag", "abcd", "help message for fstrflag")
	flag.BoolVar(&boolFlag, "fboolflag", false, "help message for fboolflag")
	flag.DurationVar(&durFlag, "fdurflag", 3*time.Minute, "help message for fdurflag")
}

func main() {
	//
	fmt.Println("Start...")

	// 
	flag.Parse() // Parse all flags

	// Print them
	fmt.Println("fintflag", intFlag)
	fmt.Println("fstrflag", strFlag)
	fmt.Println("fboolflag", boolFlag)
	fmt.Println("fdurflag", durFlag)

	//
	fmt.Println("End...")
}
