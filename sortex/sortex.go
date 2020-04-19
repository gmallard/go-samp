/*
An example of using sort.
*/

package main

import (
	"fmt"
	"sort"
)

type tt struct {
	tk int
	td string
}

var tvs = []tt{
	tt{6, "xyz1"},
	tt{2, "xyz4"},
	tt{19, "xyz3"},
	tt{7, "xyz2"},
}

// Type Stringer
func (v tt) String() string {
	return fmt.Sprintf("%d: %s", v.tk, v.td)
}

// ByTT stuff.
type ByTT []tt

func (t ByTT) Len() int {
	return len(t)
}

func (t ByTT) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t ByTT) Less(i, j int) bool {
	return t[i].tk < t[j].tk
}

// ByTD stuff.
type ByTD []tt

func (t ByTD) Len() int {
	return len(t)
}

func (t ByTD) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t ByTD) Less(i, j int) bool {
	return t[i].td < t[j].td
}

func main() {
	s := "======================================================"
	sv := tvs
	//
	fmt.Println()
	fmt.Println(s)
	fmt.Println(tvs)
	//
	fmt.Println()
	fmt.Println(s)
	sort.Sort(ByTT(tvs))
	fmt.Println(tvs)
	//
	fmt.Println()
	fmt.Println(s)
	sort.Sort(ByTD(tvs))
	fmt.Println(tvs)
	//
	fmt.Println()
	fmt.Println(s)
	fmt.Println("Orig:", sv)
	sort.Slice(sv, func(i, j int) bool {
		return sv[i].tk > sv[j].tk // Note: > ==> Descending
	})
	fmt.Println("Desc:", sv)
}
