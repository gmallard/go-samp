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

func randBetween2(min, max int64) int64 {
	rt, _ := rand.Int(rand.Reader, big.NewInt(max-min)) // Ignore errors here
	return min + rt.Int64()
}

func randBetween3(min, max int64, fact float64) int64 {
	rt, _ := rand.Int(rand.Reader, big.NewInt(max-min)) // Ignore errors here
	return int64(fact * float64(min+rt.Int64()))
}

func main() {
	fmt.Printf("MIN: %d\n", min)
	fmt.Printf("MAX: %d\n", max)
	//
	fmt.Println("StartT1...")
	for i := 1; i <= 100; i++ {
		fmt.Printf("Iteration %d\n", i)
		r := randBetween(min, max)
		fmt.Printf("Result: %d\n", r)
		if r < min || r > max {
			fmt.Printf("Error, unexpected results: %v, %v, %v\n", min, max, r)
			panic("unexpected results")
		}
	}
	fmt.Println("EndT1...")
	//
	fmt.Println("StartT2...")
	for i := 1; i <= 100; i++ {
		fmt.Printf("Iteration %d\n", i)
		r := randBetween2(min, max)
		fmt.Printf("Result: %d\n", r)
		if r < min || r > max {
			fmt.Printf("Error, unexpected results: %v, %v, %v\n", min, max, r)
			panic("unexpected results")
		}
	}
	fmt.Println("EndT2...")
	//
	fmt.Println("StartT3...")
	for i := 1; i <= 100; i++ {
		fmt.Printf("Iteration %d\n", i)
		r := randBetween3(min, max, 1.0)
		fmt.Printf("Result: %d\n", r)
		if r < min || r > max {
			fmt.Printf("Error, unexpected results: %v, %v, %v\n", min, max, r)
			panic("unexpected results")
		}
	}
	fmt.Println("EndT3...")
}
