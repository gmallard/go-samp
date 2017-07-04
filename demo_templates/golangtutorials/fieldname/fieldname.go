/*
Modified From: http://golangtutorials.blogspot.com/2011/06/go-templates.html
*/
package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name     string
	AgeField string // OK
}

func main() {
	p := Person{Name: "Mary", AgeField: "31"}

	t := template.New("nonexported template demo")
	t, _ = t.Parse("hello1 {{.Name}}! ")
	err := t.Execute(os.Stdout, p)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	//
	t, _ = t.Parse("Age {{.AgeField}}!\n")
	err = t.Execute(os.Stdout, p)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	//
	t, _ = t.Parse("hello2 {{.Name}}! Age {{.AgeField}}!\n")
	err = t.Execute(os.Stdout, p)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
}
