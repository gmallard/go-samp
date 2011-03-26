// Demo a wait group

package main

import (
	"fmt" //
  "time" //
  "sync" //
)

func called(id string, loops, sltime int, wg *sync.WaitGroup) {
  for i := 0; i < loops; i++ {
    wnum := i + 1
    fmt.Printf("id %s, waitnum: %d\n", id, wnum)
    err := time.Sleep(int64(sltime) * 1e9)
    if err != nil {
      // ??? 
    }
  }
  fmt.Println(id, "is done")
  wg.Done()
}

func main() {
	fmt.Println("Start...")
  wg := new(sync.WaitGroup)
  //
  wg.Add(1)
  go called("1", 6, 2, wg)
  //
  wg.Add(1)
  go called("2", 5, 1, wg)
  //
  wg.Add(1)
  go called("3", 8, 3, wg)
  //
  fmt.Println("Starting main wait")
  wg.Wait()
	fmt.Println("End...")
}
