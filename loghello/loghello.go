/*
Using the log package.
*/
package main

import (
	"fmt" // formatter
	"log" //
	"os"  //
)

func main() {
	fmt.Println("Start...")
	//	var logger = log.New(os.Stdout, nil, 1) // NOT OK
	//	var logger = log.New(os.Stdout, "LPR ", 1) // OK
	// var logger = log.New(os.Stdout, "LPR ", log.Ldate|log.Ltime|log.Lshortfile) // OK
	var logger = log.New(os.Stdout, "LPR ", log.Ldate|log.Lmicroseconds|log.Lshortfile) // OK
	logger.Println("logline")
	fmt.Println("End...")
}
