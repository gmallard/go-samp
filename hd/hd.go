/*
	Dump a file in hex.
*/
package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Flag variables
var (
	inFile   string
	offBegin int
	offEnd   int
	lineLen  int
	innerLen int
	goDump   bool
	//
	argFname string
	fileLen  = -1
	addrFlen = -1
)

// Main initialization, a convenient place to set flags up
func init() {
	flag.StringVar(&inFile, "inFile", "", "input file name")
	flag.IntVar(&offBegin, "offBegin", 0, "begin dump at offset")
	flag.IntVar(&offEnd, "offEnd", -1, "end dump at offset")
	flag.IntVar(&lineLen, "lineLen", 16, "dump line byte count")
	flag.IntVar(&innerLen, "innerLen", 4, "dump line inner area byte count")
	flag.BoolVar(&goDump, "goDump", false, "if true, use standard go dump")
}

func checkError(e error, ds string) {
	if e != nil {
		fmt.Printf("\n%s %s\n\n", ds, e)
		os.Exit(1)
	}
}

func hexDigitCount(i int) {
	addrFlen = 1
	for {
		i = i / 16
		if i == 0 {
			return
		}
		addrFlen++
	}
}

func setFileLen(f *os.File) {
	fi, err := f.Stat()
	checkError(err, "Stat Error ==>")
	fileLen = int(fi.Size())
	hexDigitCount(fileLen)
	addrFlen++
	if addrFlen < 6 {
		addrFlen = 6
	}
	// fmt.Printf("Hex Digit Count: %d\n", addrFlen)
}

func getReader() io.Reader {
	fmt.Printf("Args: %v\n", os.Args)
	fa := flag.Args()
	if len(fa) >= 1 {
		argFname = fa[0]
	}
	if inFile == "" && argFname == "" {
		return os.Stdin
	}
	if inFile != "" {
		f, err := os.OpenFile(inFile, os.O_RDONLY, 0644)
		checkError(err, "inFile Open Error ==>")
		setFileLen(f)
		return f
	}
	if argFname != "" {
		f, err := os.OpenFile(argFname, os.O_RDONLY, 0644)
		checkError(err, "argFname Open Error ==>")
		setFileLen(f)
		return f
	}
	// Never get here ......
	return nil
}

func goFormatDump(r io.Reader) {
	// Dump
	buff, err := ioutil.ReadAll(r)
	checkError(err, "ReadAll error ==>")
	fmt.Printf("%s", hex.Dump(buff))
	//
	return
}

func prtOff(o int) {
	had := fmt.Sprintf("%016x", o)
	fmt.Printf("%s  ", had[16-addrFlen:])
}

func blankBuf() []byte {
	s := strings.Repeat(" ", lineLen)
	return []byte(s)
}

func prtLeft(br int, ib []byte) {
	nol := lineLen / innerLen
	os := ""
	noff := 0
	for no := 0; no < nol; no++ {
		for ni := 0; ni < innerLen; ni++ {
			if noff < br {
				nbi := int(ib[noff])
				os = fmt.Sprintf("%s%02x", os, nbi)
			} else {
				os = os + "  " // Add two blanks here
			}
			noff++
		}
		os = os + " "
	}
	fmt.Print(os)
	fmt.Print(" * ")
}

func prtRight(ib []byte) {
	bb := blankBuf()
	for i := 0; i < lineLen; i++ {
		bb[i] = ib[i]
		if bb[i] < byte(0x20) {
			bb[i] = byte('.')
		}
	}
	fmt.Print(string(bb))
	fmt.Print(" *")
}

func main() {
	fmt.Println("DumpFile Starts....")
	flag.Parse() // Parse all flags
	// fmt.Println("Line Length:", lineLen)
	r := getReader()
	if goDump {
		goFormatDump(r)
		fmt.Println("DumpFile Ends....")
		return
	}
	//
	roff := offBegin
	for {
		ib := blankBuf()
		br, _ := io.ReadAtLeast(r, ib, lineLen)
		// fmt.Println("Read Length:", br)
		if br == 0 {
			fmt.Println()
			break
		}
		prtOff(roff)
		prtLeft(br, ib)
		prtRight(ib)
		fmt.Println()
		roff += lineLen
	}
	fmt.Println("DumpFile Ends....")
}
