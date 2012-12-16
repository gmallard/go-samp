/*
Brief slice of strings demonstration.
*/
package main

import (
	"fmt" //
)

type paired_slice []string

func (p paired_slice) Add(k, v string) paired_slice {
	r := append(p, k, v)
	return r
}

func main() {
	fmt.Println("Start...")
	ps := paired_slice{}
	fmt.Println("ps01:", ps)
	ps = ps.Add("k1", "v1")
	fmt.Println("ps02:", ps)
	ps = ps.Add("k2", "v2")
	fmt.Println("ps03:", ps)
	//
	ps = ps.Add("k3", "v3").Add("k4", "v4")
	fmt.Println("ps04:", ps)
	//
	fmt.Println("ps90:", paired_slice{}.
		Add("ak1", "av1").
		Add("ak2", "av2"))
	fmt.Println("End...")
}
