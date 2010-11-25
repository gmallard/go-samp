package main

import (
	"fmt"
	"http"
	"io"
)

//
// A primitive attempt at the exercise in the GoCourseDay2.pdf.
//
// I had difficulties starting because the method signatures used in the PDF
// examples are no longer valid.
//
// Along the way if found a bunch of other goodness in the distribution:
// a) effective_go.html (appears up to date)
// b) the codelab directory
//


type srvrinfo struct {
	wtr http.ResponseWriter
	req *http.Request
}

func showInfo(si srvrinfo) {
	fmt.Fprintf(si.wtr, "Method: |%s|<br />", si.req.Method)
	//
	fmt.Fprintf(si.wtr, "Raw URL: |%s|<br />", si.req.URL.Raw)
	fmt.Fprintf(si.wtr, "Scheme: |%s|<br />", si.req.URL.Scheme)
	fmt.Fprintf(si.wtr, "Raw Authority: |%s|<br />", si.req.URL.RawAuthority)
	fmt.Fprintf(si.wtr, "Host: |%s|<br />", si.req.URL.Host)
	fmt.Fprintf(si.wtr, "Raw Path: |%s|<br />", si.req.URL.RawPath)
	fmt.Fprintf(si.wtr, "Path: |%s|<br />", si.req.URL.Path)
	fmt.Fprintf(si.wtr, "Opaque Path: |%t|<br />", si.req.URL.OpaquePath)
	fmt.Fprintf(si.wtr, "Raw Query: |%s|<br />", si.req.URL.RawQuery)
	fmt.Fprintf(si.wtr, "Fragment: |%s|<br />", si.req.URL.Fragment)
}

func HelloServer(rwtr http.ResponseWriter, req *http.Request) {
	io.WriteString(rwtr, "Hello, World!!!!!!<br />")
	//
	var parms = srvrinfo{rwtr, req}
	showInfo(parms)
	//
	io.WriteString(rwtr, "Later ....<br />")
}

func main() {
	fmt.Println("Hello, World!!")
	http.Handle("/hello", http.HandlerFunc(HelloServer))
	http.ListenAndServe(":54321", nil)
}
