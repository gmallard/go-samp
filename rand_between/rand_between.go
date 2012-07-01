/*
Generate a random int64 number between min and max bounds.
*/
package main

import (
	"crypto/rand"
	"fmt"      //
	"math/big" //
)

var max int64 = 1e9
var min int64 = 10 * max / 100

//

func randBetween(min, max int64) int64 {
	br, _ := rand.Int(rand.Reader, big.NewInt(max-min)) // Ignore errors here
	return br.Add(big.NewInt(min), br).Int64()
}

func main() {
	fmt.Println("Start...")
	fmt.Printf("MIN: %d\n", min)
	fmt.Printf("MAX: %d\n", max)
	for i := 1; i <= 1000000; i++ {
		fmt.Printf("Iteration %d\n", i)
		r := randBetween(min, max)
		fmt.Printf("Result: %d\n", r)
		if r < min || r > max {
			fmt.Printf("Error, unexpected results: %v, %v, %v\n", min, max, r)
			panic("unexpected results")
		}
	}

	fmt.Println("End...")
}
