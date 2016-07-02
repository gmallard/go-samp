/*
Using maps of pointers.
*/
package main

import "fmt"

func main() {
	//
	type ts struct {
		tval int
	}
	//
	sp1 := new(ts)
	sp1.tval = 12345
	sp2 := new(ts)
	sp2.tval = 34567
	//
	fmt.Printf("sp1: [%v]\n", sp1)
	fmt.Printf("sp1: [%v]\n", sp2)
	//
	ms := make(map[string]*ts)
	ms["sp1"] = sp1
	ms["sp2"] = sp2
	//
	fmt.Printf("ms: [%v]\n", ms)
	//
	if v, ok := ms["sp1"]; ok {
		fmt.Println("sp1 check ok", v, v.tval)
	} else {
		fmt.Println("sp1 check fail")
	}
	//
	if v, ok := ms["sp2"]; ok {
		fmt.Println("sp2 check ok", v, v.tval)
	} else {
		fmt.Println("sp2 check fail")
	}
	//
	if v, ok := ms["sp3"]; ok {
		fmt.Println("sp3 check", v)
	} else {
		fmt.Println("sp3 check fail")
	}
	//
	fmt.Println()
}

