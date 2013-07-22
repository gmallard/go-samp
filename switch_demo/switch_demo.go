/*
A basic demonstration of the switch statement.
*/
package main

import (
	"fmt" //
)

func main() {
	fmt.Println("Start...")
	i := []int{1, 2, 3, 4, 42}
	for _, v := range i {
		switch v {
		case 1:
			fmt.Println("Is 1")
		case 2, 3:
			fmt.Println("Is 2 or 3")
		case 4:
			fmt.Println("Is 4")
		default:
			fmt.Println("Default", v)
		}
	}
	fmt.Println("End...")
}
