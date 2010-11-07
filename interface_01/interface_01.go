package main

import (
	"fmt"
	"math"
)

type Point struct { x, y float }
//
func (p *Point) Abs() float {
	x64 := float64(p.x)
	y64 := float64(p.y)
	r64 := math.Sqrt(x64*x64 + y64*y64)
	return float(r64)
}

// Point implements this.
// As does MyFloat.
// float64 does not.

type AbsInterface interface {
	Abs() float // receiver is implied
}


type MyFloat float

// This does *NOT* cpmpile ....
//
// func (f MyFloat) Abs() float {
// 	if f < 0.0 { return -f }
//	return f
// }
//

// This does.
func (f MyFloat) Abs() float {
	if f < 0.0 { return float(-f) }
	return float(f)
}

func main() {
  fmt.Println("Start .....")

	var ai AbsInterface;

	pp := new(Point);
	fmt.Println(pp.Abs()); // method call

	ai = pp; // OK: *Point has Abs()
	fmt.Println(ai.Abs()); // method call

	// ai = 7.; // compile-time err: float has no Abs()

	ai = MyFloat(-7.)	// OK, MyFloat has Abs()
	fmt.Println(ai.Abs()); // method call

	ai = &Point{ 3, 4 };
	fmt.Println(ai.Abs()); // method call

  fmt.Println("End .....")
}

