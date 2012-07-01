/*
Demonstrate type coercion with numbers.
*/
package main

// Demo basic coercion of numbers from one type to another

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Start .....")
	//
	var f64 float64 = 1.23456
	fmt.Printf("f64: %g\n", f64)
	//
	var f32 float32 = 3.45678
	fmt.Printf("f32: %g\n", f32)
	//
	var f float32 = 5.6789
	fmt.Printf("f: %g\n", f)
	//
	f64 = float64(f32)
	fmt.Printf("f64: %g\n", f64)
	//
	f32 = float32(f)
	fmt.Printf("f32: %g\n", f32)
	//
	f64 = float64(f32)
	fmt.Printf("f64: %g\n", f64)
	//
	f64 = math.Pi
	fmt.Printf("f64: %g\n", f64)
	f32 = float32(f64)
	fmt.Printf("f32: %g\n", f32)
	//
	fmt.Println("End .....")
}
