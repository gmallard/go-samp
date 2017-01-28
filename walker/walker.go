/*
   Demo filepath Walk
*/

package main

import (
	"log"
	"os"
	"path/filepath"
)

var (
	fc = 0
	dc = 0
)

func main() {
	walkbase := os.Getenv("WALKDIR") // Supply your own start subdirectoy here
	e := os.Chdir(walkbase)
	if e != nil {
		log.Fatalln("Chdir error:", e)
	}
	fdn, e := os.Getwd()
	if e != nil {
		log.Fatalln("Getwd error:", e)
	}
	log.Printf("Top Fitrvtory:%s\n", fdn)
	err := filepath.Walk(fdn, myWalker)
	if err != nil {
		log.Fatalln("MAIN ERROR:", err)
	}
	log.Println("Dir Count", dc)
	log.Println("File Count", fc)
}

/*
	Using *this* technique from the standard library, it does not seem
	possible to walk a directory structure looking for full path names.
	I can not figure out how to push/pop subdirectory names using
	this technique.
*/
func myWalker(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Printf("%s %v\n", "Walker Error: ", err)
		return nil
	}
	//
	if info.IsDir() {
		log.Println("Directory:", info.Name())
		dc++
	} else {
		log.Println("File:", info.Name())
		fc++
	}
	return nil
}
