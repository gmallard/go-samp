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
// label was originally coded as:
// defalt:
// misspelled.
// Originally a compile issued no warnings/errors for this.  Which can
// cause strange and hard to find runtime bugs.  However at go revision
// 7786, the misspelled coding takes a hard error:
//
//* switch_test.go:29: label defalt defined and not used
//* make[2]: *** [switch_test.6] Error 1
//
// Now coded correctly.
//
//*ORIG        defalt:
        default:
            fmt.Println("Dunno")
    }
    fmt.Println("End...")
}

