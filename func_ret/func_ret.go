/*
Using multiple return values.
*/
package main

import (
	"fmt"
	"math"
)

//
// As of Feb 2011, there is an additional difficulty with many of the
// examples publushed in the gocourseday# PDFs (as well as many other
// go examples).  This is:
//
// type 'float' has been totally removed from the language!
//
// ------------------------------
//
// There are a number of examples in the gocourseday#
// PDF files (supplied with the go distribution) that will not compile.
//
// One example (gocourseday1.pdf, page 45) is:
//
/* ----------------------------
func MySqrt(f float) (v float, ok bool) {
	if f >= 0 { 
		v,ok = math.Sqrt(f),true 
	} else {
		v,ok = 0,false 
	}
	return v,ok
}
------------------------------ */

//
// The general problem is the use of 'float' and the math.* methods actually
// take and return 'float64' types.  
// When the 'math' packages changed from 'float'
// to 'float64' is unknown to the author.
//
// One work around is to just change all 'float' types to 'float64'.  This
// may or may not be what you actually want!
//
func MySqrt(f float64) (v float64, ok bool) {
	if f >= 0 {
		v, ok = math.Sqrt(f), true
	} else {
		v, ok = 0, false
	}
	return v, ok
}

//
// And there are other, more or less suitable work arounds as well.
//

// Return with no values
func MySqrt2(f float64) (v float64, ok bool) {
	if f >= 0 {
		v, ok = math.Sqrt(f), true
	}
	return // no values, returns default: v, ok
}

// Return with no values
func MySqrt3(f float64) (v float64, ok bool) {
	if f < 0 {
		return
	} // error case
	return math.Sqrt(f), true
}

// ----------------
// Show function returns with multiple return values, and the use of
// 'default' returns.
// ----------------
func main() {
	fmt.Println("Start....")

	var fa float64 = 1.234567
	var rv float64
	var ok bool

	rv, ok = MySqrt(fa)
	fmt.Printf("SQR01: %g\t%t\n", rv, ok)

	fa = -1.234567
	rv, ok = MySqrt(fa)
	fmt.Printf("SQR02: %g\t%t\n", rv, ok)

	fa = 1.234567
	rv, ok = MySqrt2(fa)
	fmt.Printf("SQR201: %g\t%t\n", rv, ok)
	rv, ok = MySqrt3(fa)
	fmt.Printf("SQR301: %g\t%t\n", rv, ok)

	fmt.Println("End....")
}
