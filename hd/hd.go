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
	quiet    bool
	h        bool
	hexUpper bool
	edgeMark string
	//
	argFname string
	fileLen  = -1
	addrFlen = -1
)

// Main initialization, a convenient place to set flags up
func init() {
	flag.StringVar(&inFile, "inFile", "",
		"input file name.  Argument 0 may also be used.")
	flag.StringVar(&edgeMark, "edgeMark", "|",
		"single character at edges if the right hand side.")
	flag.IntVar(&offBegin, "offBegin", 0,
		"begin dump at file offset.")
	flag.IntVar(&offEnd, "offEnd", -1,
		"end dump at file offset.")
	flag.IntVar(&lineLen, "lineLen", 16,
		"dump line byte count.")
	flag.IntVar(&innerLen, "innerLen", 4,
		"dump line inner area byte count.")
	flag.BoolVar(&goDump, "goDump", false,
		"if true, use standard go encoding/hex/Dump.")
	flag.BoolVar(&hexUpper, "hexUpper", false,
		"if true, print upper case hex.")
	flag.BoolVar(&h, "h", false, "print usage message.")
}

func checkError(e error, ds string) {
	if e != nil {
		fmt.Printf("\n%s %s\n\n", ds, e)
		if !quiet {
			fmt.Println("DumpFile Ends, RC:", 1)
		}
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

func fileInit(fn, ed string) io.Reader {
	f, err := os.OpenFile(fn, os.O_RDONLY, 0644)
	checkError(err, ed+" Open Error ==>")
	setFileLen(f)
	if offBegin > 0 {
		_, err := f.Seek(int64(offBegin), io.SeekStart)
		checkError(err, "Seek Error ==>")
	}
	return f
}

func getReader() io.Reader {
	fa := flag.Args()
	if len(fa) >= 1 {
		argFname = fa[0]
	}
	if inFile == "" && argFname == "" {
		addrFlen = 8 // Arbitrary, file size is unknown
		return os.Stdin
	}
	if inFile != "" {
		return fileInit(inFile, "inFile")
	}
	if argFname != "" {
		return fileInit(argFname, "argFname")
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

func printOffset(o int) {
	hexFmt := "%016x"
	if hexUpper {
		hexFmt = "%016X"
	}
	had := fmt.Sprintf(hexFmt, o)
	fmt.Printf("%s  ", had[16-addrFlen:])
}

func blankBuf(l int) []byte {
	s := strings.Repeat(" ", l)
	return []byte(s)
}

func printLeftBuffer(br int, ib []byte) {
	nol := (lineLen / innerLen) + 1
	os := ""
	noff := 0
	hexFmt := "%s%02x"
	if hexUpper {
		hexFmt = "%s%02X"
	}
leftFor:
	for no := 0; no < nol; no++ {
		for ni := 0; ni < innerLen; ni++ {
			if noff < br {
				nbi := int(ib[noff])
				os = fmt.Sprintf(hexFmt, os, nbi)
			} else {
				os = os + "  " // Add two blanks here
			}
			noff++
			if noff > br {
				break leftFor
			}
		}
		os = os + " " // Blank at end of inner
	}
	fmt.Print(os)
	fmt.Print(edgeMark)
}

func printRightBuffer(br int, ib []byte) {
	bb := blankBuf(br)
	for i := 0; i < br; i++ {
		bb[i] = ib[i]
		if bb[i] < byte(0x20) {
			bb[i] = byte('.')
		}
	}
	fmt.Print(string(bb))
	fmt.Print(edgeMark)
}

func main() {
	flag.Parse() // Parse all flags
	if h {
		flag.PrintDefaults()
		return
	}
	if !quiet {
		fmt.Println("DumpFile Starts....")
	}
	// fmt.Println("Line Length:", lineLen)
	r := getReader()
	if goDump {
		goFormatDump(r)
		if !quiet {
			fmt.Println("DumpFile Ends....")
		}
		return
	}
	if offEnd > 0 && offEnd <= offBegin {
		fmt.Printf("Offset Error: offEnd(%d) must be > offBegin(%d)\n",
			offEnd, offBegin)
		if !quiet {
			fmt.Println("DumpFile Ends, RC:", 2)
		}
		os.Exit(2)
	}
	//
	roff := offBegin
	for {
		readLen := lineLen
		if offEnd > 0 && roff+readLen > offEnd {
			readLen = offEnd - roff + 1
		}
		ib := blankBuf(readLen)
		// fmt.Println("ReadLen is now:", readLen)
		br, _ := io.ReadAtLeast(r, ib, readLen)
		// fmt.Println("Actual Read Length:", br)
		if br == 0 {
			break
		}
		printOffset(roff)
		printLeftBuffer(br, ib)
		printRightBuffer(br, ib)
		roff += lineLen
		if offEnd > 0 && roff > offEnd {
			break
		}
		fmt.Println()
	}
	if !quiet {
		fmt.Println("DumpFile Ends....")
	}
}
