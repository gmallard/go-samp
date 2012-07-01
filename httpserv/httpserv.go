/*
A very simple HTTP server.
*/
package main

import (
	"fmt"
	"net/http"
)

//
//  Modified for weekly 2012-01-20
//

func HelloServer(rwtr http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rwtr, "Hello, World!!!!!!\n")
	//
	fmt.Fprintf(rwtr, "Method: |%s|\n", req.Method)
	fmt.Fprintf(rwtr, "Scheme: |%s|\n", req.URL.Scheme)
	fmt.Fprintf(rwtr, "Opaque: |%s|\n", req.URL.Opaque)
	fmt.Fprintf(rwtr, "User: |%v|\n", req.URL.User)
	fmt.Fprintf(rwtr, "Host: |%s|\n", req.URL.Host)
	fmt.Fprintf(rwtr, "Path: |%s|\n", req.URL.Path)
	fmt.Fprintf(rwtr, "Raw Query: |%s|\n", req.URL.RawQuery)
	fmt.Fprintf(rwtr, "Fragment: |%s|\n", req.URL.Fragment)

	//
	fmt.Fprintf(rwtr, "Later ....\n")
}

func main() {
	fmt.Println("Hello, World!!")
	http.Handle("/hello", http.HandlerFunc(HelloServer))
	http.ListenAndServe(":54321", nil)
}
