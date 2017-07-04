/*
Example from:
https://astaxie.gitbooks.io/build-web-application-with-golang/en/07.4.html
*/

package main

import (
	"os"
	"text/template"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	f1 := Friend{Fname: "Chris"}
	f2 := Friend{Fname: "George"}
	t := template.New("fieldname example 2")
	t, _ = t.Parse(`hello {{.UserName}}!
            {{range .Emails}}
                an email {{.}}
            {{end}}
			{{/* Next section */}}
            {{with .Friends}}
            {{range .}}
                my friend's name is {{.Fname}}
            {{end}}
            {{end}}
            `)
	p := Person{UserName: "Guy",
		Emails:  []string{"guy@gmallard.com", "allard.guy.m@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}
