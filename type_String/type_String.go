/*
An interesting demo of types and methods.
*/
package main

//
// Methods are not just for structs. They can be
// defined for any (non-pointer) type.
//
// The type must be defined in your package though.
//
// You can't write a method for int but you can
// declare a new int type and give it methods.
//
// fmt.Print?? responds to String()
//
import "fmt"

type Day int

var dayName = []string{
	"Sunday", "Monday", "Tuesday", "Wednesday",
	"Thursday", "Friday", "Saturday",
}

func (day Day) String() string {
	return dayName[day]
}

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println("Start .....")
	//
	var day = Tuesday
	fmt.Print(day.String())
	fmt.Println()
	fmt.Println(0, Sunday, 1, Monday)
	fmt.Println("End .....")
}
