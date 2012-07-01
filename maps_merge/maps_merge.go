/*
Merging two maps, with optional overwrite.
*/
package main

import "fmt"

type Mymss map[string]string

func mapMerge(ma Mymss, mb Mymss, overlay bool) {
	var present bool
	for bk, bv := range mb {
		if overlay {
			ma[bk] = bv
		} else {
			_, present = ma[bk]
			if !present {
				ma[bk] = bv
			}
		}
	}
	return
}

func main() {
	fmt.Println("Start...")
	//
	var mss1 = Mymss{"a": "av", "b": "bv", "c": "cv"}
	fmt.Println("mss1:")
	for k, v := range mss1 {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}
	//
	var mss2 = Mymss{"d": "dv", "e": "ev"}
	fmt.Println("mss2:")
	for k, v := range mss2 {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}
	//
	mapMerge(mss1, mss2, false)
	fmt.Println("merged1:")
	for k, v := range mss1 {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}

	//
	var mss3 = Mymss{"c": "cnew", "d": "dv", "e": "ev"}
	fmt.Println("mss3:")
	for k, v := range mss3 {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}
	//
	mss1 = Mymss{"a": "av", "b": "bv", "c": "cv"} // Reset
	mapMerge(mss1, mss3, false)
	fmt.Println("merged2:")
	for k, v := range mss1 {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}
	//
	mss1 = Mymss{"a": "av", "b": "bv", "c": "cv"} // Reset
	mapMerge(mss1, mss3, true)
	fmt.Println("merged3:")
	for k, v := range mss1 {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}
	//
	fmt.Println("End...")
}
