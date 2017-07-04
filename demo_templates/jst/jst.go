/*
js function example.
*/
package main

import (
	"os"
	"text/template"
)

type Person struct {
	UserName string
	UserID   int
}

func main() {
	// Not sure I see a practical use for the 'js' function.
	t := template.New("JS example")
	as := `{{js .}}
`
	t, _ = t.Parse(as)
	p := Person{UserName: "Guy", UserID: 123456}
	// The output here is somewhat unexpected.  At the very least, it is not
	// what I anticipated.
	t.Execute(os.Stdout, p)
}
