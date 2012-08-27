/*
Simple template example.
*/
package main

import (
	"fmt"
	"os"
	"text/template"
)

type Data struct {
	Title string
	Vi    int
	Vs    string
}

const td = `
This is the title: {{.Title}}

The int is: {{.Vi}}
The string is: {{.Vs}}
`

var Datum = Data{Title: "Chapter 1", Vi: 123, Vs: "bogus line"}

func main() {
	fmt.Println("Datum:", Datum)
	//
	t1 := template.New("T1")
	t1, err := t1.Parse(td)
	if err != nil {
		fmt.Println("err1", err)
	}
	err = t1.Execute(os.Stdout, Datum)
	if err != nil {
		fmt.Println("err2", err)
	}
}
