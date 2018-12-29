package main

/*
Show executable file name and directory.
*/
import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	pgm, e := os.Executable()
	if e != nil {
		log.Fatalln("Error:", e)
	}
	fmt.Printf("Executable Name:\n[%s]\n", pgm)
	//
	dir := filepath.Dir(pgm)
	fmt.Printf("Executable Directory:\n[%s]\n", dir)
}
