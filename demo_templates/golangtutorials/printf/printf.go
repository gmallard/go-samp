/*
From:
http://golangtutorials.blogspot.com/2011/10/go-templates-part-2.html
*/
package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.New("test")
	t = template.Must(t.Parse("{{with $x := `hello`}}{{printf `%s %s` $x `Mary`}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)
}
