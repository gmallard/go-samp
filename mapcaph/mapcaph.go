/*
Map capacity hint demo.
*/
package main

import "fmt"

func main() {
	mm := make(map[string]int, 2) // 2 is a "capacity hint", not a hard bound
	/* From the language spec:
	The initial capacity does not bound its size: maps grow to accommodate the
	number of items stored in them, with the exception of nil maps. A nil map
	is equivalent to an empty map except that no elements may be added.
	*/

	// Also note that maps do not support the 'cap' builtin.  Saying:
	//     cap(mm)
	// is an compile error, with a message like:
	// invalid argument mm (type map[string]int) for cap

	mm["a"] = 1
	fmt.Println(mm["a"], len(mm))
	mm["b"] = 2
	fmt.Println(mm["b"], len(mm))
	mm["c"] = 3
	fmt.Println(mm["c"], len(mm))
	mm["d"] = 4
	fmt.Println(mm["d"], len(mm))
}
