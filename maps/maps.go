/*
Using maps, from the gocourse PDFs.
*/
package main

import "fmt"

func main() {
	//
	var value float32
	var present bool
	//
	// Create a map using a literal.  Access a particular vlaue using the key.
	//
	var m = map[string]float32{"1": 1, "pi": 3.1415}
	fmt.Printf("%g\n", m["pi"])
	//
	// Add a key and value.
	//
	m["2"] = 2
	fmt.Printf("%g\n", m["2"])
	//
	// Change a value!
	//
	m["2"] = 3
	fmt.Printf("%g\n", m["2"])
	//
	// Determine if a key is present in the map.
	//
	value, present = m["1"]                // value is there
	fmt.Printf("%g\t%t\n", value, present) // present=true
	value, present = m["6"]                // value is not there
	fmt.Printf("%g\t%t\n", value, present) // present=false
	//
	// Show how not using the ", OK" format works.
	//
	fmt.Println()
	crash := m["nothere"]     // Examples said crash.  Different I suppose.
	fmt.Printf("%g\n", crash) // Gives zero value for type
	//
	// Loop over all map members/elements.
	//
	fmt.Println()
	for key, value := range m {
		fmt.Printf("%s\t%g\n", key, value)
	}
	fmt.Println()
	//
	// Loop over just the keys.
	//
	fmt.Println()
	for key := range m {
		fmt.Printf("%s\n", key)
	}
	fmt.Println()
	//
	// Delete a element/member.
	//
	value, present = m["2"]                // value is there
	fmt.Printf("%g\t%t\n", value, present) // present=true
	delete(m, "2")                         // Delete, use the zero value for the values, and false
	value, present = m["2"]                // value is not there
	fmt.Printf("%g\t%t\n", value, present) // present=false
	fmt.Println()
}
