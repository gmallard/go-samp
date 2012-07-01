/*
Interface example from the gocourse PDFs.
*/
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float32
}

//
func (p *Point) Abs() float32 {
	x64 := float64(p.x)
	y64 := float64(p.y)
	r64 := math.Sqrt(x64*x64 + y64*y64)
	return float32(r64)
}

// Point implements this.
// As does MyFloat.
// float64 does not.

type AbsInterface interface {
	Abs() float32 // receiver is implied
}

type MyFloat float32

// This does *NOT* cpmpile ....
// Even when 'float' was supported ....
//
// func (f MyFloat) Abs() float {
// 	if f < 0.0 { return -f }
//	return f
// }
//

// This does.
func (f MyFloat) Abs() float32 {
	if f < 0.0 {
		return float32(-f)
	}
	return float32(f)
}

func main() {
	fmt.Println("Start .....")

	var ai AbsInterface

	pp := new(Point)
	fmt.Println(pp.Abs()) // method call

	ai = pp               // OK: *Point has Abs()
	fmt.Println(ai.Abs()) // method call

	ai = MyFloat(-7.)     // OK, MyFloat has Abs()
	fmt.Println(ai.Abs()) // method call

	ai = &Point{3, 4}
	fmt.Println(ai.Abs()) // method call

	fmt.Println("End .....")
}
