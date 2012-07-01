/*
Demonstrate using the List container, embellished.
*/
package main

import (
	"container/list" //
	"fmt"            //
)

type Pair struct {
	key   string
	value string
}

func generate(key, value string) Pair {
	var pair Pair
	pair.key = key
	pair.value = value
	return pair
}

func lexamp01() {
	var myl = list.New()
	fmt.Printf("myl: %v\n", myl)
	//
	var pa Pair = generate("key1", "value1")
	fmt.Printf("Key: %s, Value: %s\n", pa.key, pa.value)
	myl.PushBack(pa)
	//
	pa = generate("key2", "value2")
	myl.PushBack(pa)
	//
	pa = generate("key3", "value3")
	myl.PushBack(pa)
	//
	fmt.Printf("myl: %v\n", myl)
	//
	fmt.Println("==================================")
	for nelt := myl.Front(); nelt != nil; nelt = nelt.Next() {
		fmt.Printf("Next: %v\n", nelt)
		if np, ok := nelt.Value.(Pair); ok {
			fmt.Println("ok", np)
			fmt.Printf("Next Key: %s, Next Value: %s\n", np.key, np.value)
		}
	}
	//
	fmt.Println("==================================")
	for nelt := myl.Front(); nelt != nil; nelt = nelt.Next() {
		fmt.Printf("Next: %v\n", nelt)
		np, _ := nelt.Value.(Pair)
		fmt.Printf("Next Key: %s, Next Value: %s\n", np.key, np.value)
	}
}

func main() {
	fmt.Println("Start...")
	lexamp01()
	fmt.Println("End...")
}
