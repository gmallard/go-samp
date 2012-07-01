/*
Demonstrate all the tricks from http://code.google.com/p/go-wiki/wiki/SliceTricks
*/
package main

// http://code.google.com/p/go-wiki/wiki/SliceTricks

import (
	"fmt"
)

var a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var b = []int{10, 11, 12, 13, 14}

// Print a slice of int's.
func prtsl_int(a []int) {
	fmt.Printf("len: %d, cap: %d\n", len(a), cap(a))
	fmt.Println(a)
	fmt.Println()
}

// Reset at start of each demo.
func reset() {
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b = []int{10, 11, 12, 13, 14}
}

// AppendVector
func showAppend(o string) {
	fmt.Println(o)
	reset()
	a = append(a, b...)
	fmt.Println("===app1===")
	prtsl_int(a)
}

// Copy
func showCopy(o string) {
	fmt.Println(o)
	reset()
	copy(a, b) // Overlays
	fmt.Println("===cop1===")
	prtsl_int(a)

	t := make([]int, len(b))
	reset()
	copy(t, a) // Surprise maybe, truncates
	fmt.Println("===cop2===")
	prtsl_int(t)

}

// Cut
func showCut(o string) {
	fmt.Println(o)
	reset()
	a = append(a[:3], a[6:]...)
	fmt.Println("===cut1===")
	prtsl_int(a)
}

// Delete
func showDel(o string) {
	fmt.Println(o)
	reset()
	i := 4
	a = append(a[:i], a[i+1:]...)
	fmt.Println("===del1===")
	prtsl_int(a)
}

// Expand
func showExpand(o string) {
	fmt.Println(o)
	reset()
	i := 3
	j := 5
	a = append(a[:i], append(make([]int, j), a[i:]...)...)
	prtsl_int(a)
}

// Extend
func showExtend(o string) {
	fmt.Println(o)
	reset()
	j := 3
	a = append(a, make([]int, j)...)
	prtsl_int(a)
}

// Insert
func showInsert(o string) {
	fmt.Println(o)
	reset()
	i := 2
	a = append(a[:i], append([]int{42}, a[i:]...)...)
	prtsl_int(a)
}

// InsertVector
func showInsertVector(o string) {
	fmt.Println(o)
	reset()
	i := 3
	a = append(a[:i], append(b, a[i:]...)...)
	prtsl_int(a)
}

// Pop
func showPop(o string) {
	fmt.Println(o)
	reset()
	x, a := a[len(a)-1], a[:len(a)-1]
	prtsl_int(a)
	fmt.Println("x", x)
}

// Push
func showPush(o string) {
	fmt.Println(o)
	reset()
	a = append(a, 42)
	prtsl_int(a)
}

func main() {

	fmt.Println("===a===")
	prtsl_int(a)

	fmt.Println("===b===")
	prtsl_int(b)

	showAppend("***Append***")
	showCopy("***Copy***")
	showCut("***Cut***")
	showDel("***Delete***")
	showExpand("***Expand***")
	showExtend("***Extend***")
	showInsert("***Insert***")
	showInsertVector("***InsertVector***")
	showPop("***Pop***")
	showPush("***Push***")
}
