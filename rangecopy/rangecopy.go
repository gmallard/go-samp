/*
Demonstrate that range makes copies.
*/
package main

import (
	"fmt"
)

type Foo struct {
	a int
}

//
// Based on discussion:
// https://groups.google.com/forum/?fromgroups=#!topic/golang-nuts/e08r1Vk7ufQ
//
func main() {

	// A range of primitives
	someInts := []int{1, 2}
	fmt.Println("somei1", someInts[0], someInts[1])

	for _, si := range someInts {
		si = si + 1
	}
	fmt.Println("somei2", someInts[0], someInts[1]) // Not modified

	for _, si := range someInts {
		z := &si
		*z = *z + 1
	}
	fmt.Println("somei3", someInts[0], someInts[1]) // Not modified

	for i := range someInts { // A slight surprize
		someInts[i] = someInts[i] + 1
	}
	fmt.Println("somei3B", someInts[0], someInts[1]) // Modified

	for i := 0; i < len(someInts); i++ {
		someInts[i] = someInts[i] + 1
	}
	fmt.Println("somei4", someInts[0], someInts[1]) // Modified

	for i := 0; i < len(someInts); i++ {
		z := &someInts[i]
		*z = *z + 1
	}
	fmt.Println("somei5", someInts[0], someInts[1]) // Modified again

	afoo := Foo{1}
	fmt.Println("afooa", afoo.a)

	// A range of structs
	someFoos := []Foo{Foo{2}, Foo{3}}
	fmt.Println("somex1", someFoos[0].a, someFoos[1].a)

	for _, sf := range someFoos {
		sf.a = sf.a + 1
	}
	fmt.Println("somex2", someFoos[0].a, someFoos[1].a) // Not modified

	for _, sf := range someFoos {
		z := &sf
		z.a = z.a + 1
	}
	fmt.Println("somex3", someFoos[0].a, someFoos[1].a) // Not modified

	for i := range someFoos {
		someFoos[i].a = someFoos[i].a + 1
	}
	fmt.Println("somex3B", someFoos[0].a, someFoos[1].a) // Modified

	for i := 0; i < len(someFoos); i++ {
		someFoos[i].a = someFoos[i].a + 1
	}
	fmt.Println("somex4", someFoos[0].a, someFoos[1].a) // Modified

	for i := 0; i < len(someFoos); i++ {
		z := &someFoos[i]
		z.a = z.a + 1
	}
	fmt.Println("somex5", someFoos[0].a, someFoos[1].a) // Modified again

}
