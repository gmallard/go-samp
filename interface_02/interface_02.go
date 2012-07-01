/*
Interface example from gocourseday2.pdf.
*/
package main

import (
	"fmt"
)

type MyFloat float32

//
// The gocourseday2.pdf shows this function definition as:
// func (f MyFloat) Abs() float { ...
// which does not compile.
// After some thinking, I guess this is what is (might be?) intended:
//
func (f MyFloat) Abs() MyFloat {
	if f < 0.0 {
		return -f
	}
	return f
}

//
// And then the experiments with variable 'ftryb' below suggested this 
// approach/technique.  Perhaps it is what is intended in the PDF.
//
func (f MyFloat) Abs2() float32 {
	if f < 0.0 {
		return float32(-f)
	}
	return float32(f)
}

//
func main() {
	fmt.Println("Start .....")
	var mfa MyFloat
	fmt.Printf("%v\n", mfa)

	// OK, assign ...
	mfa = 1.234
	fmt.Printf("%v\n", mfa)

	// Change value and do the Abs() method
	mfa = -3.14159
	fmt.Printf("Neg: %v\n", mfa)
	fmt.Printf("Abs: %v\n", mfa.Abs())

	// Sanity check
	var ftrya float32
	fmt.Printf("TryA: %v\n", ftrya)

	// Coerced conversion
	var ftryb float32 = 0.0
	ftryb = float32(mfa)
	fmt.Printf("TryB: %v\n", ftryb)

	// And use if alternate Abs2() method
	var mfb MyFloat = -4.567
	fmt.Printf("MFBa: %v\n", mfb)
	fmt.Printf("MFBaAbs2: %v\n", mfb.Abs2())

	// And yeah, we can add MyFloat types .....
	var mfc MyFloat
	mfc = mfa + mfb
	fmt.Printf("MFC: %v\n", mfc)

	//
	fmt.Println("End .....")
}
