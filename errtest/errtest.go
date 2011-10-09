// Fiddling with Error types .....

package main

import (
	"fmt" //
	"os"
)

type Error string

func (e Error) String() (string) {
	return string(e)
}

const (
	ERRA = Error("error A")
	ERRB = Error("error B")
)

func badA() (os.Error) {
	return ERRA
}

func badB() (os.Error) {
	return ERRB
}

func main() {
	fmt.Println("Start...")

	ea := badA()
	fmt.Println("ea:", ea)

	eb := badB()
	fmt.Println("eb:", eb)

	if ea == ERRA {
		fmt.Println("yep, ea is OK")
	}

	if ea != eb {
		fmt.Println("yep, eb is OK also")
	}

	fmt.Println("End...")
}
