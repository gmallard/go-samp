/*
   Demo a 'os' package based directory walker.
*/
package main

import (
	//"fmt"
	"log"
	"os"
	"path"
)

var (
	seltype = ""
)

func walkDir(wd string) {
	fd, e := os.Open(wd)
	if e != nil {
		log.Fatalln("Open error:", e)
	}
	log.Printf("===== Start Next Dir:%s\n", wd)
	var subdirs []string
	fil, e := fd.Readdir(-1)
	if e != nil {
		log.Fatalln("Readdir error:", e)
	}
	// Process Files
	for _, nfi := range fil {
		prtFileInfo(nfi)
		if nfi.IsDir() {
			subdirs = append(subdirs, nfi.Name())
			continue
		}
		nfn := wd + string(os.PathSeparator) + nfi.Name()
		log.Printf("Next File Name:%s\n", nfn)
		if seltype != "" {
			fext := path.Ext(nfn)
			log.Printf("File Extension:%s\n", fext)
			if fext != "" {
				fext = fext[1:] // Eliminate the '.'
			}
			if fext == seltype {
				log.Printf("Selected File:%s\n", nfn)
			}
		}
	}
	// Process subdirs
	for _, nsd := range subdirs {
		absd := wd + string(os.PathSeparator) + nsd
		walkDir(absd)
	}
}

func main() {
	log.Printf("Start")
	seltype = os.Getenv("SELTYPE") // Use this to 'select' a file extension
	log.Printf("Select Type:%s\n", seltype)
	walkbase := os.Getenv("WALKDIR") // Supply your own start subdirectoy here
	e := os.Chdir(walkbase)
	if e != nil {
		log.Fatalln("Chdir error:", e)
	}
	fdn, e := os.Getwd()
	if e != nil {
		log.Fatalln("Getwd error:", e)
	}
	log.Println("Full Path:", fdn)
	walkDir(fdn)
	log.Printf("End")
}

func prtFileInfo(fi os.FileInfo) {
	log.Println("===== Next File Info:")
	log.Printf("Name:%s\n", fi.Name())
	log.Printf("Size:%d\n", fi.Size())
	fm := fi.Mode()
	bp := fm.Perm()
	log.Printf("Mode:%v (%0o)\n", fm, uint32(bp))
	log.Printf("ModTime:%v\n", fi.ModTime())
	log.Printf("IsDir:%v\n", fi.IsDir())
	//log.Printf("Sys:%v\n", fi.Sys())
}
