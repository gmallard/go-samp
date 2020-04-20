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
	tt{19, "xyz7"},
	tt{19, "xyz3"},
	tt{7, "xyz2"},
}

var (
	sv = tvs
)

// Type Stringer
func (v tt) String() string {
	return fmt.Sprintf("%d->%s", v.tk, v.td)
}

// ByTK stuff.
type ByTK []tt

func (t ByTK) Len() int {
	return len(t)
}

func (t ByTK) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t ByTK) Less(i, j int) bool {
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
	//
	fmt.Println()
	fmt.Println(s)
	fmt.Println("Original")
	fmt.Println(tvs)
	//
	fmt.Println()
	fmt.Println(s)
	sort.Sort(ByTK(tvs))
	fmt.Println("By Key")
	fmt.Println(tvs)
	//
	fmt.Println()
	fmt.Println(s)
	sort.Sort(ByTD(tvs))
	fmt.Println("By Data")
	fmt.Println(tvs)
	//
	fmt.Println()
	fmt.Println(s)
	fmt.Println("Orig:", sv)
	sort.Slice(sv, func(i, j int) bool {
		return sv[i].tk > sv[j].tk // Note: > ==> Descending
	})
	fmt.Println("Desc By Key:", sv)
}
