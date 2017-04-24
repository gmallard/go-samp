/*
Demo slices based on the discussion here:

https://blog.golang.org/slices

*/
package main

import (
	"bytes"
	"fmt"
)

/*
Imagine a slice header is defined like this:

    type sliceHeader struct {
        Length        int
        Capacity      int
        ZerothElement *byte // or *whatever type
    }
*/

var (
	buffer   [256]byte
	breakLen = 10
)

type path []byte

func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i]
	}
}

func (p path) ValTruncateAtFinalSlash() {
	i := bytes.LastIndex(p, []byte("/"))
	if i >= 0 {
		p = p[0:i]
	}
}

// This can work because the value receiver still points at the same
// array.
func (p path) ToUpper() {
	for i, b := range p {
		if 'a' <= b && b <= 'z' {
			p[i] = b + 'A' - 'a'
		}
	}
}

func fillBuffer() {
	for i := range buffer {
		buffer[i] = byte(i)
	}
}

func slinfo(s []byte) {
	fmt.Println("Address: ", &s[0], "LEN", len(s), "CAP", cap(s))
	ppsl(s)
	fmt.Println()
}

func slinfoInt(s []int) {
	fmt.Println("Address: ", &s[0], "LEN", len(s), "CAP", cap(s))
	ppslInt(s)
	fmt.Println()
}

func ppsl(s []byte) {
	// fmt.Println("Address: ", &s[0])
	for i := range s {
		fmt.Printf("0x%02x ", s[i])
		if (i+1)%breakLen == 0 || i == (len(s)-1) {
			fmt.Println()
		}
	}
}

func ppslInt(s []int) {
	// fmt.Println("Address: ", &s[0])
	for i := range s {
		fmt.Printf("%d ", s[i])
		if (i+1)%breakLen == 0 || i == (len(s)-1) {
			fmt.Println()
		}
	}
}

func SubtractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]
	return slice
}

func PtrSubtractOneFromLength(slicePtr *[]byte) {
	slice := *slicePtr
	*slicePtr = slice[0 : len(slice)-1]
}

func Extend(slice []int, element int) []int {
	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func demoExtend() {
	var iBuffer [10]int
	islice := iBuffer[0:0]
	for i := 0; i < 20; i++ {
		islice = Extend(islice, i)
		// fmt.Println(islice)
		slinfoInt(islice)
	}
}

func doubleCap() {
	slice := make([]int, 10, 15)
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
	newSlice := make([]int, len(slice), 2*cap(slice))
	for i := range slice {
		newSlice[i] = slice[i]
	}
	slice = newSlice
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
}

func doubleCapCopy() {
	slice := make([]int, 10, 15)
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
	newSlice := make([]int, len(slice), 2*cap(slice))
	/*
	  The copy function is smart. It only copies what it can, paying attention
	  to the lengths of both arguments. In other words, the number of elements it
	  copies is the *minimum* of the lengths of the two slices. This can save a
	  little bookkeeping. Also, copy returns an integer value, the number of
	  elements it copied, although it's not always worth checking.
	*/
	ecc := copy(newSlice, slice)
	fmt.Println("copied", ecc, "elements")
	slice = newSlice
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
}

// Insert inserts the value into the slice at the specified index,
// which must be in range.
// The slice must have room for the new element.
func Insert(slice []int, index, value int) []int {
	// Grow the slice by one element.
	slice = slice[0 : len(slice)+1]
	// Use copy to move the upper part of the slice out of the way and open a hole.
	// (That this works is not intuitive to me at least .........)
	ecc := copy(slice[index+1:], slice[index:])
	fmt.Println("copied", ecc, "elements")
	// Store the new value.
	slice[index] = value
	// Return the result.
	return slice
}

func demoInsert() {
	slice := make([]int, 10, 20) // Note capacity > length: room to add element.
	for i := range slice {
		slice[i] = i
	}
	fmt.Println("insert1", slice, "len", len(slice), "cap", cap(slice))
	slice = Insert(slice, 5, 99)
	fmt.Println("insert2", slice, "len", len(slice), "cap", cap(slice))
}

func ExtendNoPanic(slice []int, element int) []int {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]int, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func demoExtendNoPanic() {
	slice := make([]int, 0, 5)
	for i := 0; i < 10; i++ {
		slice = ExtendNoPanic(slice, i)
		fmt.Printf("len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
		fmt.Println("address of 0th element:", &slice[0])
	}
}

// Append appends the items to the slice.
// First version: just loop calling Extend (the no panic version).
func Append(slice []int, items ...int) []int {
	for _, item := range items {
		slice = ExtendNoPanic(slice, item)
	}
	return slice
}

func demoAppend() {
	slice := []int{0, 1, 2, 3, 4}
	fmt.Println("append1A", slice)
	slice = Append(slice, 5, 6, 7, 8)
	fmt.Println("append1B", slice)
}

func demoAppend2() {
	slice1 := []int{0, 1, 2, 3, 4}
	slice2 := []int{55, 66, 77}
	fmt.Println("append2A", slice1)
	slice1 = Append(slice1, slice2...) // The '...' is essential!
	// The '...' "explodes" the slice value into individual elements
	fmt.Println("append2B", slice1)
}

// Append appends the elements to the slice.
// Efficient version.
func AppendEfficient(slice []int, elements ...int) []int {
	n := len(slice)
	total := len(slice) + len(elements)
	if total > cap(slice) {
		// Reallocate. Grow to 1.5 times the new size, so we can still grow.
		newSize := total*3/2 + 1
		newSlice := make([]int, total, newSize)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:total]
	copy(slice[n:], elements)
	return slice
}

func demoAppendEfficient() {
	slice1 := []int{0, 1, 2, 3, 4}
	slice2 := []int{55, 66, 77}
	fmt.Println("appendefficient1A", slice1)
	slice1 = AppendEfficient(slice1, slice2...) // The '...' is essential!
	fmt.Println("appendefficient1B", slice1)
}

func demoAppendBI() {
	// Demo the append builtin
	// Create a couple of starter slices.
	slice := []int{1, 2, 3}
	slice2 := []int{55, 66, 77}
	fmt.Println("Start slice: ", slice)
	fmt.Println("Start slice2:", slice2)

	// Add an item to a slice.
	slice = append(slice, 4)
	fmt.Println("Add one item:", slice)

	// Add one slice to another.
	slice = append(slice, slice2...)
	fmt.Println("Add one slice:", slice)

	// Make a copy of a slice (of int).
	slice3 := append([]int(nil), slice...)
	fmt.Println("Copy a slice:", slice3)

	// Copy a slice to the end of itself.
	fmt.Println("Before append to self:", slice)
	slice = append(slice, slice...)
	fmt.Println("After append to self:", slice)

}

func main() {
	fillBuffer()
	allslice := buffer[:]
	fmt.Println("allslice:")
	slinfo(allslice)
	//
	slice := buffer[10:20]
	fmt.Println("slice:")
	slinfo(slice)
	for i := 0; i < len(slice); i++ {
		slice[i] = slice[i] + 1
	}
	fmt.Println("slice before:")
	slinfo(slice)
	newslice := SubtractOneFromLength(slice)
	fmt.Println("slice after:")
	slinfo(slice)
	fmt.Println("newslice:")
	slinfo(newslice)
	//
	PtrSubtractOneFromLength(&slice)
	fmt.Println("slice after ptrsub:")
	slinfo(slice)

	// It is idiomatic to use a pointer receiver for a method that
	// modifies a slice.

	fmt.Println()
	pathName := path("/usr/bin/tso") // Conversion from string to path.
	fmt.Printf("pathName1a:[%s]\n", pathName)
	pathName.TruncateAtFinalSlash()
	fmt.Printf("pathName1b:[%s]\n", pathName)

	pathName = path("/usr/bin/tso")
	fmt.Printf("pathName2a:[%s]\n", pathName)
	pathName.ValTruncateAtFinalSlash() // This basically does not work
	fmt.Printf("pathName2b:[%s]\n", pathName)

	pathName = path("/usr/bin/tso")
	fmt.Printf("pathName3a:[%s]\n", pathName)
	pathName.ToUpper() // This will work
	fmt.Printf("pathName3b:[%s]\n", pathName)

	fmt.Println("==========================================")
	fmt.Println("capacity demos")

	// capacity discussion
	// demoExtend() // This panics, which is the point of the demo

	/*
	   What if we want to grow the slice beyond its capacity? You can't! By
	   definition, the capacity is the limit to growth. But you can achieve
	   an equivalent result by allocating a new array, copying the data over,
	   and modifying the slice to describe the new array.
	*/
	doubleCap()     // Demo growing a slice
	doubleCapCopy() // Demo growing a slice another way
	//
	demoInsert() // Demo insert an element

	// append
	demoExtendNoPanic() // This works

	//
	demoAppend()          // This does as well
	demoAppend2()         // This does as well
	demoAppendEfficient() // And this

	/*
	  And so we arrive at the motivation for the design of the append built-in
	  function. It does exactly what our AppendEfficient example does, with equivalent
	  efficiency, but it works for any slice type.
	*/
	demoAppendBI()
}
