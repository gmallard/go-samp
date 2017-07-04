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
	t, _ := template.New("test").Parse("{{with `hello`}}{{.}}{{end}}!\n")
	t.Execute(os.Stdout, nil)

	t1, _ := template.New("test").Parse("{{with `hello`}}{{.}} {{with `Mary`}}{{.}}{{end}}{{end}}!\n") //when nested, the dot takes the value according to closest scope.
	t1.Execute(os.Stdout, nil)
}
