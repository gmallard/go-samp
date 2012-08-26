/*
Reading XML.  An initial demonstration.
*/
package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Flags
var (
	fileName string
)

type Area51 struct {
	Location string `xml:"location,attr"`
	Password string `xml:"password,attr"`
	Data     string `xml:",chardata"`
}

type Configs struct {
	XMLName xml.Name `xml:"configs"`
	Car54   string   `xml:"car54"`
	Area51  Area51   `xml:"area51"`
}

var buff []byte // File data

// Main initialization, a convenient place to set flags up
func init() {
	flag.StringVar(&fileName, "fileName", "./data.xml", "xml file name") // This example uses the default
}

/*
Loads file contents to the 'buff' slice.
*/
func loadFile() {
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("\nOpen Error => %s\n\n", err)
		os.Exit(1)
	}

	// Load the entire file to the 'buff' slice
	reader := bufio.NewReader(f) // Buffered reader
	buff, err = ioutil.ReadAll(reader)
	if err != nil {
		fmt.Printf("\nReadAll Error => %s\n\n", err)
		os.Exit(1)
	}

	// Close
	err = f.Close()
	if err != nil {
		fmt.Printf("\nClose Error => %s\n\n", err)
		os.Exit(1)
	}
}

func showXml() {
	ctx := Configs{}
	err := xml.Unmarshal(buff, &ctx)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("Configs: %+v\n", ctx)
	// The main element name
	fmt.Printf("XMLName: %#v\n", ctx.XMLName)
	// Trim, because there will usually be leading and maybe trailing whitespace
	fmt.Printf("Car54: %q\n", strings.TrimSpace(ctx.Car54))
	fmt.Printf("Location: %q\n", ctx.Area51.Location)
	fmt.Printf("Password: %q\n", ctx.Area51.Password)
	fmt.Printf("Area51: %q\n", strings.TrimSpace(ctx.Area51.Data))
}

func main() {
	// Hi!
	fmt.Println("Start...")
	// Flag handling.
	flag.Parse() // Parse all flags
	fmt.Println("fileName", fileName)

	// Load the entire file
	loadFile()

	// Display struct values from the XML
	showXml()

	// Bye!
	fmt.Println("End...")
}
