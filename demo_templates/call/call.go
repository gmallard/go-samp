/*
From:
https://github.com/golang/go/issues/10653
https://play.golang.org/p/Ka8bN3_V1f
*/
package main

import (
	"os"
	"text/template"
)

type TestStructA struct {
	Test func(string) string
}

func test(s string) string {
	return "Test A Says " + s + "\n"
}

type TestStructB struct {
}

func (self TestStructB) Test() func(string) string {
	return func(s string) string { return "Test B Says " + s + "\n" }
}

// DOES NOT WORK
type TestStructC struct {
}

func (self TestStructC) Test(s string) string {
	return "Test C Says " + s + "\n"
}

func noCall() {
	tempStr := "{{.Test \"hi\"}}" // NB: no 'call' function used
	t := template.Must(template.New("test_temp2").Parse(tempStr))
	testC := TestStructC{}
	err := t.Execute(os.Stdout, testC)
	if err != nil {
		panic(err)
	}
}

func main() {
	tempStr := "{{call .Test \"hi\"}}"
	t := template.Must(template.New("test_temp").Parse(tempStr))

	testA := TestStructA{Test: test}
	err := t.Execute(os.Stdout, testA)
	if err != nil {
		panic(err)
	}

	testB := TestStructB{}
	err = t.Execute(os.Stdout, testB)
	if err != nil {
		panic(err)
	}

	noCall() // other way to do the 'does not work' part below

	// DOES NOT WORK
	testC := TestStructC{}
	err = t.Execute(os.Stdout, testC)
	if err != nil {
		panic(err)
	}
}
