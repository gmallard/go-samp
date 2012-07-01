/*
Demonstrate using the List container.
*/
package main

import (
	"container/list" //
	"fmt"            //
)

func lexamp01() {
	var myl = list.New()
	if myl == nil {
		fmt.Println("is nil")
	}
	fmt.Printf("Len: %v\n", myl.Len())
	fmt.Printf("myl: %v\n", myl)
	var elts = []string{"elt1", "elt2", "elt3"}
	for _, nelt := range elts {
		myl.PushBack(nelt)
	}
	fmt.Printf("myl: %v\n", myl)
	//
	for nelt := myl.Front(); nelt != nil; nelt = nelt.Next() {
		fmt.Printf("Next: %v\n", nelt.Value)
	}
}

func main() {
	fmt.Println("Start...")
	lexamp01()
	fmt.Println("End...")
}
