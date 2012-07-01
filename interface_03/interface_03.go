/*
Interface example, using interface{}.
*/
package main

import (
	"fmt"
	"reflect"
)

type hmap map[string]string
type hstr []string

func pickem(iv interface{}) {
	switch iv.(type) {
	case hmap:
		fmt.Println("is hmap")
	case hstr:
		fmt.Println("is hstr")
	default:
		fmt.Println("NOT!")
	}
}

func dumpem(iv interface{}) {
	fmt.Println("raw", iv)
	fmt.Println("type:", reflect.TypeOf(iv))
	v := reflect.ValueOf(iv)
	fmt.Println("value:", v)
	fmt.Println("kind:", v.Kind())
	fmt.Println("interface:", v.Interface())
}

func showdata(iv interface{}) {
	switch w := iv.(type) {
	case hmap:
		fmt.Println("showdata hmap")
		for k, v := range w {
			fmt.Println("key:", k, "value:", v)
		}
	case hstr:
		fmt.Println("showdata hstr")
		for i, v := range w {
			fmt.Println("index:", i, "value:", v)
		}
	default:
		fmt.Println("SHOWDATA NOT!")
	}

}

func main() {
	fmt.Println("Start .....")
	//
	h := hmap{"a": "va"}
	sa := hstr{"a", "b"}
	//
	fmt.Println("==========================")
	pickem(h)
	fmt.Println("--------------------------")
	dumpem(h)
	fmt.Println("..........................")
	showdata(h)
	fmt.Println("==========================")
	pickem(sa)
	fmt.Println("--------------------------")
	dumpem(sa)
	fmt.Println("..........................")
	showdata(sa)
	//
	fmt.Println("End .....")
}
