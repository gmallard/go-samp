/*
Using OS environment variables.
*/
package main

import (
	"fmt" //
	"os"  //
)

func main() {
	fmt.Println("Start...")
	//
	env_hello := os.Getenv("PATH") // Probably here
	fmt.Printf("PATH is: !%s!\n", env_hello)
	//
	env_hello = os.Getenv("SOME_JUNK") // Probably not here
	fmt.Printf("SOME_JUNK is: !%s!\n", env_hello)
	//
	fmt.Println("End...")
}
