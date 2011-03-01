// Switch Test
package main
import (
    "fmt" //
)
//
// A very unintuitive gitcha.
// This compiles, and produces incorrect results.
// It probably should compile - the 'dafaut' is taken as a label.
// Difficult to catch this during testing.
//
func main() {
    fmt.Println("Start...")
    anint := 256
    switch(anint) {
        case 1:
            fmt.Println("Is 1")
        case 2:
            fmt.Println("Is 2")
        case 3:
            fmt.Println("Is 3")
        defalt:
            fmt.Println("Dunno")
    }
    fmt.Println("End...")
}

