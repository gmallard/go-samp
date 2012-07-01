/*
Fiddling with Error types.
*/
package main

import "fmt" //

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ERRA = Error("error A")
	ERRB = Error("error B")
)

func badA() error {
	return ERRA
}

func badB() error {
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
