/*
From:
http://golangtutorials.blogspot.com/2011/10/go-templates-part-2.html
*/
package main

/*
You can create local variables for the pipelines within the template by
prefixing the variable name with a "$" sign. Variable names have to be composed
of alphanumeric characters and the underscore. In the example below I have used
a few variations that work for variable names.
*/
import (
	"os"
	"text/template"
)

func main() {
	t := template.Must(template.New("name").Parse("{{with $3 := `hello`}}{{$3}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)

	t1 := template.Must(template.New("name1").Parse("{{with $x3 := `hola`}}{{$x3}}{{end}}!\n"))
	t1.Execute(os.Stdout, nil)

	t2 := template.Must(template.New("name2").Parse("{{with $x_1 := `hey`}}{{$x_1}} {{.}} {{$x_1}}{{end}}!\n"))
	t2.Execute(os.Stdout, nil)
}
