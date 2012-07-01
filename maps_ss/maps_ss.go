/*
Simple play with map[string]string.
*/
package main

import "fmt"

type Mymss map[string]string

func main() {
	//
	fmt.Println("Start...")
	var mss1 = map[string]string{"a": "av", "b": "bv"}
	for k, v := range mss1 {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}
	//
	var mss2 = Mymss{"a2": "av2", "b2": "bv2"}
	for k, v := range mss2 {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}
	//
	mss3 := Mymss{"a3": "av3", "b3": "bv3"}
	for k, v := range mss3 {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}
	fmt.Println("End...")
}
