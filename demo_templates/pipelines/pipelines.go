/*
Templates - from the golang package documentation examples.
*/
package main

import (
	"fmt"
	"os"
	"text/template" // *not* html/template
)

func main() {
	fmt.Println("start")
	pipedemo1()
	fmt.Println()
	pipedemo2()
	fmt.Println()
	pipedemo3()
	fmt.Println()
	pipedemo4()
	fmt.Println()
	pipedemo5()
	fmt.Println()
	pipedemo6()
	fmt.Println()
	pipedemo7()
	fmt.Println()
	pipedemo8()
	fmt.Println()
	pipedemo9()
	fmt.Println()
	pipedemo10()
	fmt.Println()
	pipedemo11()
	fmt.Println("end")
}

var pipel1 = `{{"\"output1\""}}
`

func pipedemo1() {
	t1 := template.New("T1")
	t1, err := t1.Parse(pipel1)
	if err != nil {
		fmt.Println("err1", err)
	}
	err = t1.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("err2", err)
	}
}

// This example is a bit different than what is shown in the godoc .......
// Break this text up and embellish it in order to get clean compile,
// execute, and produce correct(?) output
var pipel2beg = "{{"
var pipel2main = `"\"output2\"\n"`
var pipel2end = "}}"
var pipel2all = pipel2beg + pipel2main + pipel2end

func pipedemo2() {
	t1 := template.New("T1")
	t1, err := t1.Parse(pipel2all)
	if err != nil {
		fmt.Println("err1", err)
	}
	err = t1.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("err2", err)
	}
}

var pipel3 = `{{printf "%q" "output3"}}
`

func pipedemo3() {
	t1 := template.New("T1")
	t1, err := t1.Parse(pipel3)
	if err != nil {
		fmt.Println("err1", err)
	}
	err = t1.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("err2", err)
	}
}

var pipel4 = `{{"output4" | printf "%q"}}
`

func pipedemo4() {
	t1 := template.New("T1")
	t1, err := t1.Parse(pipel4)
	if err != nil {
		fmt.Println("err1", err)
	}
	err = t1.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("err2", err)
	}
}

var pipel5 = `{{printf "%q" (print "out" "put5")}}
`

func pipedemo5() {
	t1 := template.New("T1")
	t1, err := t1.Parse(pipel5)
	if err != nil {
		fmt.Println("err1", err)
	}
	err = t1.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("err2", err)
	}
}

var pipel6 = `{{"put6" | printf "%s%s" "out" | printf "%q"}}
`

func pipedemo6() {
	t1 := template.New("T1")
	t1, err := t1.Parse(pipel6)
	if err != nil {
		fmt.Println("err1", err)
	}
	err = t1.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("err2", err)
	}
}

var pipel7 = `{{"output7" | printf "%s" | printf "%q"}}
`

func pipedemo7() {
	t1 := template.New("T1")
	t1, err := t1.Parse(pipel7)
	if err != nil {
		fmt.Println("err1", err)
	}
	err = t1.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("err2", err)
	}
}

var pipel8 = `{{with "output8"}}{{printf "%q" .}}{{end}}
`

func pipedemo8() {
	t1 := template.New("T1")
	t1, err := t1.Parse(pipel8)
	if err != nil {
		fmt.Println("err1", err)
	}
	err = t1.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("err2", err)
	}
}

var pipel9 = `{{with $x := "output9" | printf "%q"}}{{$x}}{{end}}
`

func pipedemo9() {
	t1 := template.New("T1")
	t1, err := t1.Parse(pipel9)
	if err != nil {
		fmt.Println("err1", err)
	}
	err = t1.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("err2", err)
	}
}

var pipel10 = `{{with $x := "output10"}}{{printf "%q" $x}}{{end}}
`

func pipedemo10() {
	t1 := template.New("T1")
	t1, err := t1.Parse(pipel10)
	if err != nil {
		fmt.Println("err1", err)
	}
	err = t1.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("err2", err)
	}
}

var pipel11 = `{{with $x := "output11"}}{{$x | printf "%q"}}{{end}}
`

func pipedemo11() {
	t1 := template.New("T1")
	t1, err := t1.Parse(pipel11)
	if err != nil {
		fmt.Println("err1", err)
	}
	err = t1.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("err2", err)
	}
}
