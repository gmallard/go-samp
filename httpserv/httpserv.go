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

func HelloServer(rwtr http.ResponseWriter, req *http.Request) {
	io.WriteString(rwtr, "Hello, World!!!!!!<br />")
	fmt.Fprintf(rwtr, "Method: %s<br />", req.Method)
	io.WriteString(rwtr, "Later ....<br />")
}

func main() {
  fmt.Println("Hello, World!!")
	http.Handle("/hello", http.HandlerFunc(HelloServer))
	http.ListenAndServe(":54321", nil)
}

