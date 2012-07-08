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

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := First("golang", fakeSearch("replica 1"), fakeSearch("replica 2"))
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println("elapsed:", elapsed)
}
