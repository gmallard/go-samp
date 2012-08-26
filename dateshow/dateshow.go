/*
Show oddities about date formatting.
*/
package main

import (
	"fmt"
	"time"
)

// Date formatting may show seemingly incorrect values when using a 
// local Location.
func main() {
	fmt.Println("Start...")
	//
	l := "01/02/06"                // Layout string
	d := "08/26/12"                // Date string
	n, err := time.Parse(l, d[:8]) // Get time.Time
	if err != nil {
		panic(err)
	}
	fmt.Println("time.Time1:", n) // UTC, mm/dd/yy is correct
	//
	i := n.Unix()                // Seconds
	nd := time.Unix(int64(i), 0) // Has local Location, here at present: EDT
	// For this example may show 08/25/12.
	fmt.Println("time.Time2:", nd) // Local, e.g. EST, mm/dd/yy may be incorrect
	//
	nu := nd.UTC()                 // Force to UTC
	fmt.Println("time.Time3:", nu) // UTC again, mm/dd/yy is correct
	//
	fmt.Println("End...")
}
