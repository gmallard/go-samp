package main

import "fmt"

func main() {
  var m = map[string]float { "1":1, "pi":3.1415 }
	fmt.Printf("%g\n", m["pi"])
	//
	m["2"] = 2
	fmt.Printf("%g\n", m["2"])
	//
	m["2"] = 3
	fmt.Printf("%g\n", m["2"])
	//
	var value float
	var present bool
	value, present = m["1"]
	fmt.Printf("%g\t%t\n", value, present)	// present=true
	value, present = m["6"]
	fmt.Printf("%g\t%t\n", value, present)	// present=false
	//
	fmt.Println()
	for key, value := range m {
		fmt.Printf("%s\t%g\n", key, value)
	}
	//
	fmt.Println()
	for key := range m {
		fmt.Printf("%s\n", key)
	}
	//
	fmt.Println()
	crash := m["nothere"]
	fmt.Printf("%g\n", crash)	// Gives zero value for type
}

