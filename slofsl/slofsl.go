/*
Demonstrate using a slice of slices.

This example uses an anonymous struct for the demonstration.  This is
like table driven tests are frequently accomplished in the go standard library.
*/
package main

import (
	"fmt"
)

// The data for this demonstration
var sosd = []struct {
	flag int        // A flag for the test
	sos  [][]string // The slice of slices
}{
	{
		100,
		[][]string{[]string{"a100"}},
	},
	{
		200,
		[][]string{[]string{"a200", "b200"}},
	},
	{
		300,
		[][]string{[]string{"a300", "b300", "c300"}},
	},
	{
		400,
		[][]string{[]string{"a400", "b400", "d300"},
			[]string{"x400", "y400", "z300"}},
	},
}

var (
	sl = "==========================================================="
	sa = "--------------------------"
)

func main() {
	fmt.Println("start")
	fmt.Println(sl)

	// Process the slices of the anonymous struct
	for oi, ov := range sosd {
		fmt.Println()
		fmt.Println(wrap("outer"))
		fmt.Printf("oi:%d ov%q\n", oi, ov)
		fmt.Printf("struct flag:%d\n", ov.flag)

		// Process the slice of slices of strings
		for mi, mv := range ov.sos {
			fmt.Println(wrap("middle"))
			fmt.Printf("\tmi:%d mv%q\n", mi, mv)

			// Process the innermost slice
			for ii, iv := range mv {
				fmt.Println(wrap("inner"))
				fmt.Printf("\t\tii:%d iv:%q\n", ii, iv)
			}

		}

	}

	fmt.Println(sl)
	fmt.Println("end")
}

func wrap(s string) string {
	return sa + " " + s + " " + sa
}
