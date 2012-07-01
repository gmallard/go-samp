/*
A basic demonstration of the switch statement.
*/
package main

import (
	"fmt" //
)

func main() {
	fmt.Println("Start...")
	anint := 256
	switch anint {
	case 1:
		fmt.Println("Is 1")
	case 2:
		fmt.Println("Is 2")
	case 3:
		fmt.Println("Is 3")
	default:
		fmt.Println("Dunno")
	}
	fmt.Println("End...")
}
