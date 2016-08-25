/*
   Demo chaining
*/
package main

import (
	"fmt"
)

type chainer interface {
	FuncA(a int) *chainer
	FuncB(b int) *chainer
	FuncC(c int) *chainer
}

type cdata struct {
	a, b, c int
}

func (d *cdata) getA() int {
	return d.a
}

func (d *cdata) getB() int {
	return d.b
}

func (d *cdata) getC() int {
	return d.c
}

func (d *cdata) FuncA(a int) *cdata {
	d.a = a
	return d
}

func (d *cdata) FuncB(b int) *cdata {
	d.b = b
	return d
}

func (d *cdata) FuncC(c int) *cdata {
	d.c = c
	return d
}

func showData(d cdata) {
	fmt.Printf("a:%d\n", d.a)
	fmt.Printf("b:%d\n", d.b)
	fmt.Printf("c:%d\n", d.c)
}

func main() {
	fmt.Println("hi")
	//
	md := new(cdata)
	showData(*md)
	md.FuncA(1).FuncB(2).FuncC(3)
	showData(*md)

	fmt.Println("bye")

}
