/*
A more realistic demonstration of interfaces.
*/
package main

import (
	"fmt" //
)

type Gap interface {
	// Get data
	Get() (s string, e error)
	// Put data
	Put(s string) (e error)
}

// Will implement Gap
type ImpOne struct {
	ios string
}

// Implement Gap.Get
func (o *ImpOne) Get() (s string, e error) {
	return o.ios, nil
}

// Implement Gap.Put
func (o *ImpOne) Put(s string) (e error) {
	o.ios = s
	return nil
}

// Will also implement Gap
type ImpTwo struct {
	its string
}

// Implement Gap.Get
func (o *ImpTwo) Get() (s string, e error) {
	return o.its + o.its, nil
}

// Implement Gap.Put
func (o *ImpTwo) Put(s string) (e error) {
	o.its = s + s
	return nil
}

func main() {
	fmt.Println("Start...")

	// ImpOne pointer
	dOne := &ImpOne{"abcd"}
	fmt.Println("dOne.ios:", dOne.ios)

	// ImpTwo pointer
	dTwo := &ImpTwo{"1234"}
	fmt.Println("dTwo.its:", dTwo.its)

	// ImpOne leaves data alone
	_ = dOne.Put("1time")
	w, _ := dOne.Get()
	fmt.Println("dOnePutGet", "1time", w)

	// ImpTwo munges the data
	_ = dTwo.Put("1time")
	w, _ = dTwo.Get()
	fmt.Println("dTwoPutGet", "1time", w)

	// An array of Gaps
	ga := [...]Gap{dOne, dTwo}

	// Use array offsets
	w, _ = ga[0].Get()
	fmt.Println("ga0Get", w)
	w, _ = ga[1].Get()
	fmt.Println("ga1Get", w)

	// Declare a Gap and use offset 0
	var g Gap = ga[0]
	w, _ = g.Get()
	fmt.Println("gga0Get", w)

	// Now use offset 1
	g = ga[1]
	w, _ = g.Get()
	fmt.Println("gga1Get", w)

	fmt.Println("End...")
}
