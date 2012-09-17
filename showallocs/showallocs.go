/*
Show memory statistics gathering.
*/
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	stats := make([]runtime.MemStats, 60)
	for i := range stats {
		runtime.ReadMemStats(&stats[i])
		time.Sleep(time.Second)
	}
	for _, m := range stats {
		fmt.Printf("%d, %d, %d, %d, %d\n", runtime.NumGoroutine(), m.Alloc,
			m.TotalAlloc, m.HeapAlloc, m.NumGC)
	}
}
