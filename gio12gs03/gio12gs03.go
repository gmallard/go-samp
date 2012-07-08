/*
Example code based on Google I/O 2012 concurrency presentation.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string

type Search func(query string) Result

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google(query string) (results []Result) {

	// Code from the presentation:
	c := make(chan Result)
	// Run queries concurrently (fanIn)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()
	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println("elapsed:", elapsed)
}
