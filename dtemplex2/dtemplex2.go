/*
Template example.
*/
package main

import (
	"fmt"
	"os"
	"text/template"
)

type Headers []string

type Body []uint8

type Data struct {
	Command string
	Headers Headers
	Body    Body
}

func (h Headers) String() string {
	s := ""
	for i := 0; i < len(h); i += 2 {
		s = s + h[i] + "," + h[i+1] + "\n"
	}
	return s
}

func (b Body) String() string {
	return string(b)
}

const td = `{{.Command}}
{{.Headers}}
{{.Body}}
`

var Datum = Data{Command: "ACOMMAND",
	Headers: Headers{"K1", "v1", "K2", "v2"},
	Body:    []uint8("test")}

func main() {
	// fmt.Println("Datum:", Datum)
	// fmt.Println("thp:", Datum.Headers)
	// fmt.Println("thb:", Datum.Body)
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
