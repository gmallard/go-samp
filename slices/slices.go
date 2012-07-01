/*
Slice demonstration from the gocourse PDFs.
*/
package main

import "fmt"

/*
From examples.
*/
//
// Print a slice of int's.
//
func prtsl_int(a []int) {
	fmt.Printf("len: %d, cap: %d\n", len(a), cap(a))
	fmt.Println(a)
	fmt.Println()
}

//
func appendToSlice(i int, sl []int) []int {
	//
	// Exceeding capacity should be handled, but is not:
	// if len(sl) == cap(sl) { error(...) }
	//
	n := len(sl)     // Current len
	sl = sl[0 : n+1] // extend length by 1
	sl[n] = i        // store caller's value
	return sl
}

//
func main() {
	//
	// Slice literals look like an array literal without a size.
	//
	var slice = []int{1, 2, 3, 5, 6, 7, 9, 10, 11}
	prtsl_int(slice)
	//
	sla := slice[:]
	prtsl_int(sla)
	//
	slb := slice[4:]
	prtsl_int(slb)
	//
	slc := slice[4:6]
	prtsl_int(slc)
	// Range is over 'len', not 'cap'
	for _, nse := range slc {
		fmt.Printf("Next: %d\n", nse)
	}
	fmt.Println()
	//
	var s25 = make([]int, 25)
	prtsl_int(s25)
	//
	for n := 0; n < cap(s25); n++ {
		// fmt.Printf("Next: %d\n", n)
		s25[n] = n
	}
	prtsl_int(s25)
	for _, n := range s25 {
		fmt.Printf("Next: %d\n", n)
	}
	//
	fmt.Println()
	var ar = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // An array, not a slice
	fmt.Println(ar)
	//
	var sld = ar[3:5] //
	fmt.Println(sld)
	//
	var sl = make([]int, 0, 100) // len 0, cap 100
	prtsl_int(sl)
	sl = appendToSlice(123, sl)
	prtsl_int(sl)
	sl = appendToSlice(456, sl)
	prtsl_int(sl)
}
