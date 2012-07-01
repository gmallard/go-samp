/*
A demonstration of methods.
*/
package main

//
import (
	"fmt"
	"math"
)

//
// Methods can be attached to (almost?) any type.
// Declared, separatey from the type, as methods with an explicit receiver.
//
// Rules:
//
// Methods are attached to a named type, say Foo, and
// are statically bound.
//
// The type of a receiver in a method can be either *Foo
// or Foo. You can have some Foo methods and some
// *Foo methods.
//
// Foo itself cannot be a pointer type, although the
// methods can have receiver type *Foo.
//
// The type Foo must be defined in the same package as
// all its methods.
//
type Point struct {
	x, y float64
}

//
func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

//
var pointa = Point{3.0, 4.0}
var pp *Point = &pointa // OK, but not necessary.  See example code.
//
// A method does not require a pointer as a receiver.  But this is more 
// expensive:  the Point3 is passed by value.
//
type Point3 struct {
	x, y, z float64
}

func (p Point3) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

var pointb = Point3{3.0, 4.0, 5.0}

//
// A non-struct example.
//
type IntVector []int

func (v IntVector) Sum() (s int) {
	//	for i, x := range v {	// Does *NOT* compile
	for _, x := range v {
		s += x
	}
	return
}

//
// When an anonymous field is embedded in
// a struct, the methods of that (anon field) type are embedded as
// well - in effect, the struct inherits the methods of the anon field
// type.
//
type NamedPoint struct {
	Point
	name string
}

// method override
type NamedPoint2 struct {
	Point
	name string
}

func (n *NamedPoint2) Abs() float64 {
	return n.Point.Abs() * 100.
}

//
func main() {
	fmt.Println("Start....")

	// Pass a pointer to type:
	fmt.Printf("Res1: %g\n", pp.Abs())

	// Pass type by value:
	fmt.Printf("Res3: %g\n", pointb.Abs())

	// Non-struct:
	fmt.Printf("ResNS: %d\n", IntVector{1, 2, 3}.Sum())

	// go will automatically indirect/dereference values when invoking methods.
	p1 := Point{3.0, 4.0}               // p1 is a value
	fmt.Printf("Res1A: %g\n", p1.Abs()) // Abs takes a pointer ...
	p2 := &Point3{3.0, 4.0, 5.0}        // p2 is a pointer
	fmt.Printf("Res3A: %g\n", p2.Abs()) // Abs takes a value ...

	// anonymous field 'method inheritance'
	n := &NamedPoint{Point{3, 4}, "Pythagoras"}
	fmt.Println("ResANON:", n.Abs()) // prints 5

	// anonymous field 'method override'
	nb := &NamedPoint2{Point{3, 4}, "Pythagoras"}
	fmt.Println("ResANON2:", nb.Abs()) // prints 500

	fmt.Println("End....")
}
