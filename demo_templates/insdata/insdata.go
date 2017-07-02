/*
From: https://astaxie.gitbooks.io/build-web-application-with-golang/en/07.4.html
*/
package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}!\n")
	p := Person{UserName: "Guy"}
	t.Execute(os.Stdout, p)
}
