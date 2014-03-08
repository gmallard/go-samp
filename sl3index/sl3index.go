/*
Demo 3 index slicing.  Requires go 1.2.
*/
package main

import (
	"fmt"
)

func main() {
	// Init
	var array [10]int
	for i := 0; i < 10; i++ {
		array[i] = i + 1
	}

	fmt.Printf("array: %v\n", array)
	slice := array[2:4]
	fmt.Printf("slice: %v %d %d\n", slice, len(slice), cap(slice))
	fmt.Printf("array: %v\n", array)
	slice[1] = 77
	fmt.Printf("slice: %v %d %d\n", slice, len(slice), cap(slice))
	fmt.Printf("array: %v\n", array)
	slice2 := array[2:4:6]
	fmt.Printf("slice2: %v %d %d\n", slice2, len(slice2), cap(slice2))
	slice2 = append(slice2, 43)
	fmt.Printf("slice2: %v %d %d\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("array: %v\n", array)
	slice2[2] = 99
	fmt.Printf("slice2: %v %d %d\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("array: %v\n", array)
	slice2 = append(slice2, 88)
	fmt.Printf("slice2: %v %d %d\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("array: %v\n", array)
	for i := 101; i <= 105; i++ {
		slice2 = append(slice2, i)
		fmt.Printf("slice2: %v %d %d\n", slice2, len(slice2), cap(slice2))
		fmt.Printf("array: %v\n", array)
	}
	// Output:
	/*
	   array: [1 2 3 4 5 6 7 8 9 10]
	   slice: [3 4] 2 8
	   array: [1 2 3 4 5 6 7 8 9 10]
	   slice: [3 77] 2 8
	   array: [1 2 3 77 5 6 7 8 9 10]
	   slice2: [3 77] 2 4
	   slice2: [3 77 43] 3 4
	   array: [1 2 3 77 43 6 7 8 9 10]
	   slice2: [3 77 99] 3 4
	   array: [1 2 3 77 99 6 7 8 9 10]
	   slice2: [3 77 99 88] 4 4
	   array: [1 2 3 77 99 88 7 8 9 10]
	   slice2: [3 77 99 88 101] 5 8
	   array: [1 2 3 77 99 88 7 8 9 10]
	   slice2: [3 77 99 88 101 102] 6 8
	   array: [1 2 3 77 99 88 7 8 9 10]
	   slice2: [3 77 99 88 101 102 103] 7 8
	   array: [1 2 3 77 99 88 7 8 9 10]
	   slice2: [3 77 99 88 101 102 103 104] 8 8
	   array: [1 2 3 77 99 88 7 8 9 10]
	   slice2: [3 77 99 88 101 102 103 104 105] 9 16
	   array: [1 2 3 77 99 88 7 8 9 10]
	*/

}
