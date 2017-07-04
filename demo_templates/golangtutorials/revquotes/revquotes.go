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
	t := template.New("template test")
	t = template.Must(t.Parse("This is just static text. \n{{\"This is pipeline data - because it is evaluated within the double braces.\"}} {{`So is this, but within reverse quotes.`}}\n"))
	t.Execute(os.Stdout, nil)
}
