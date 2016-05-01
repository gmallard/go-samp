package main

import (
	"archive/zip"
	"os"
	// "fmt"
	"log"
	"time"
)

var (
	pref = "zip_test "
	zfl  = []string{"testdata/testfile.zip",
		"testdata/testa.jar",
	}

	alog *log.Logger
)

func init() {
	alog = log.New(os.Stdout, "ZFLG ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
}

func main() {
	// Open a zip archive for reading.
	for _, zfn := range zfl {
		r, err := zip.OpenReader(zfn)
		if err != nil {
			alog.Fatal(err)
		}
		defer r.Close()
		alog.Println(pref, "Starts", "Compressed File:", zfn)
		alog.Println()
		// Iterate through the files in the archive,
		// printing some of their contents.
		for _, f := range r.File {
			alog.Println(pref, "======Next File/Directory======")
			fhp := &f.FileHeader
			fi := fhp.FileInfo()
			dumpFileData(f, fi)
			//alog.Println("zip_test: ", "Name:", f.Name)
			// alog.Printf("Contents of %s:\n", f.Name)
			/*
			   rc, err := f.Open()
			   if err != nil {
			       alog.Fatal(err)
			   }
			   _, err = io.CopyN(os.Stdout, rc, 68)
			   if err != nil {
			       alog.Fatal(err)
			   }
			   rc.Close()
			*/
		}
		alog.Println()
		alog.Println(pref, "Done", "Compressed File:", zfn)
		alog.Println()
	}
}

func dumpFileData(zf *zip.File, zi os.FileInfo) {
	alog.Println(pref, "Name:", zf.Name, "Type:", ftype(zi.IsDir()))
	//	alog.Println(pref, "Size:", zi.Size(), "(bytes),", "Modtime:",
	//		zi.ModTime().Format(time.RFC3339))
	//	alog.Println(pref, "Mode:", zi.Mode())
	alog.Println(pref, "Mode:", zi.Mode(), "Size:", zi.Size(), "(bytes),",
		"Modtime:", zi.ModTime().Format(time.RFC3339))
}

func ftype(b bool) string {
	if b {
		return "directory"
	}
	return "file"
}
