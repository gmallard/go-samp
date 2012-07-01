/*
Demonstration of structs, based on the gocourse PDFs.
*/
package main

import "fmt"

// Structs: semple declaration of data fields.  Like C.  Memory layout.

// basic section:
// A variable is a struct
var point struct {
	x, y float32
}

// More usual
type Point struct {
	x, y float32
}

var pa Point

// Or a pointer
var paPtr *Point = new(Point)

// anons section:
//
type A struct {
	ax, ay int
}
type B struct {
	A
	bx, by float32
}

//
type C struct {
	x float32
	int
	string
}

// conflicts section:
//
type CCSA struct {
	a int
}
type CCSB struct {
	a, b int
}
type CCSC struct {
	CCSA
	CCSB
}

var c CCSC

type CCSD struct {
	CCSB
	b float32
}

var d CCSD

// literals section
//
var la = CCSA{123}
var lb = CCSB{456, 789}

// Note: full definition of inners structs is required
var lc = CCSC{CCSA{4}, CCSB{5, 6}}
var ld = CCSD{CCSB{6, 7}, 1.618}

func basics() {
	//
	fmt.Printf("01: x=%g, y=%g\n", point.x, point.y)
	point.x = 1
	fmt.Printf("02: x=%g, y=%g\n", point.x, point.y)
	//
	fmt.Printf("03: x=%g, y=%g\n", pa.x, pa.y)
	pa.x = 1
	fmt.Printf("04: x=%g, y=%g\n", pa.x, pa.y)
	//
	// Do *NOT* do this:
	// *paPtr = pa
	// It will compile and yield unexpected (to me) results.  It almost seems
	// to do a copy ???

	// Do this:
	paPtr = &pa

	fmt.Printf("05: x=%g, y=%g\n", paPtr.x, paPtr.y)
	paPtr.y = 2

	fmt.Printf("06: x=%g, y=%g\n", paPtr.x, paPtr.y)
	fmt.Printf("07: x=%g, y=%g\n", pa.x, pa.y)

	// Sanity check:
	pa.x = 123
	pa.y = 234
	paPtr.x = 345
	paPtr.y = 456
	fmt.Printf("08: x=%g, y=%g\n", paPtr.x, paPtr.y)
	fmt.Printf("09: x=%g, y=%g\n", pa.x, pa.y)
	//
}

func anons() {
	// Acts like B has four fields
	// Note: literals need detail
	b := B{A{1, 2}, 3.0, 4.0}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
	// Access by type name .....
	c := C{3.5, 7, "hello"}
	fmt.Println(c.x, c.int, c.string)
}

func conflicts() {
	//
	fmt.Println(c)
	fmt.Printf("CB: %d\n", c.b) // c.b has no conflict
	// fmt.Printf("CA: %d\n", c.a)	// Compile error, ambiguous
	fmt.Printf("cCA: %d\n", c.CCSA.a) // OK
	fmt.Printf("cBA: %d\n", c.CCSB.a) // OK
	//
	fmt.Printf("DB: %g\n", d.b)       // d.b has no conflict, it's the float32
	fmt.Printf("DBB: %d\n", d.CCSB.b) // d.CCSB.b has no conflict
}

func literals() {
	//
	fmt.Println(la)
	fmt.Printf("la.a: %d\n", la.a)
	//
	fmt.Println(lb)
	fmt.Printf("lb.a: %d lb.b%d\n", lb.a, lb.b)
	//
	fmt.Println(lc)
	fmt.Printf("lc.A.a: %d\n", lc.CCSA.a)
	fmt.Printf("lc.B.a: %d\n", lc.CCSB.a)
	fmt.Printf("lc.b: %d\n", lc.b)
	//
	fmt.Println(ld)
	fmt.Printf("ld.B.a: %d\n", ld.CCSB.a)
	fmt.Printf("ld.B.b: %d\n", ld.CCSB.b)
	fmt.Printf("ld.b: %g\n", ld.b)
}

//
func main() {
	fmt.Println("Start .....")
	basics()
	anons()
	conflicts()
	literals()
	fmt.Println("End .....")
}
