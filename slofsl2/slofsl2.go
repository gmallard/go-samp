package main

/*
Another experiment with a slice of slices, as well
as a template experiment.
*/
import (
	"fmt"
	"html/template"
	"os"
)

// table data
type tdata struct {
	Colh  []string
	Rdata [][]string
}

// template to generate HTML table
// this template uses the trick:
// {{/*
// */}}...
//
// to eliminate empty lines.  Ugly, but it works.
var tp = `
<table>
<tr>{{/*
*/}}{{range .Colh}}
<th>{{.}}</th>{{/*
*/}}{{end}}
</tr>{{/*
*/}}{{range .Rdata}}
<tr>{{/*
*/}}{{range .}}
<td>{{.}}</td>{{/*
*/}}{{end}}
</tr>{{/*
*/}}{{end}}
</table>
`

// generate an HTML table from the data
func tbgen(tdv tdata) {
	t := template.Must(template.New("").Parse(tp))
	t.Execute(os.Stdout, tdv)
}

// build the data
func exdata() tdata {
	tdv := tdata{}
	tdv.Colh = make([]string, 0)
	tdv.Rdata = make([][]string, 0)
	//
	tdv.Colh = append(tdv.Colh, "Cha", "Chb", "Chc")
	//
	ws := make([]string, 0)
	ws = append(ws, "A", "B", "C")
	tdv.Rdata = append(tdv.Rdata, ws)
	//
	ws = make([]string, 0)
	ws = append(ws, "D", "E", "F")
	tdv.Rdata = append(tdv.Rdata, ws)
	//
	ws = make([]string, 0)
	ws = append(ws, "G", "H", "I")
	tdv.Rdata = append(tdv.Rdata, ws)
	return tdv
}

// basic data dump
func dumptd(tdv tdata) {
	//
	for i, d := range tdv.Colh {
		fmt.Printf("Col: %d header is: %s\n", i, d)
	}
	fmt.Println()
	for i, d := range tdv.Rdata {
		fmt.Printf("Row: %d is: %s\n", i, d)
		for j, sd := range d {
			fmt.Printf("\tCol: %d is: %s\n", j, sd)
		}
	}
}

func main() {
	fmt.Println("hi")
	rhd := exdata()
	fmt.Printf("rhd: %v\n", rhd)
	dumptd(rhd)
	tbgen(rhd)
	//
}
