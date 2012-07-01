/*
Go tour, number 44.
*/
package main

// go tour #44: http://tour.golang.org/#44

import (
	"fmt"
	"math"
)

const (
	d = 0.0000000001 // Quit looking when this close
	m = 100          // Max loops, to prevent infinite loop
	n = 20           // Maximim square root to caclulate
)

func Sqrt(x float64) float64 {
	r := x // First guess
	for i := 1; i <= m; i++ {
		l := r                // Previous result
		r = r - (r*r-x)/(2*r) // New result
		// fmt.Println("l:", l, "r:", r)
		if (l - r) < d { // Within tolerance?
			break
		}
	}
	return r
}

func main() {
	// Estimage square roots using Newton's method
	for i := 1; i <= n; i++ {
		f := float64(i)
		fmt.Println(i, Sqrt(f), math.Sqrt(f))
	}
}
