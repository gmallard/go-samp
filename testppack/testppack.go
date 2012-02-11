package main
//
import (
	"fmt"
	"go-samp/numbers"
)
/*
	Note: numbers is a locally built package, not part of the 'go' 
	distribution.
*/
func main() {
	v := numbers.Double(6)
	fmt.Printf("6 doubled: %d\n", v)
}
