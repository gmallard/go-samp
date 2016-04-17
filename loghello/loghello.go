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
	var logger1 = log.New(os.Stdout, "LPR1 ", log.Ldate|log.Lmicroseconds|log.Lshortfile) // OK
	var logger2 = log.New(os.Stdout, "LPR2 ", log.Ldate|log.Lmicroseconds|log.Lshortfile) // OK

	logger1.Println("logline L1a")
	logger2.Println("logline L2a")
	logger1.Println("logline L1b")
	logger2.Println("logline L2b")

	fmt.Println("End...")
}
