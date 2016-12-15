/*
Display PEM file information.
*/
package main

import (
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"os"
)

var file_name = "../certs/server/server.crt"

func main() {
	fmt.Println("start.....")
	fi, e := os.Lstat(file_name)
	if e != nil {
		log.Fatalln("Lstat Error::", e.Error())
	}
	fmt.Printf("File Size:\t%d\n", fi.Size())
	//
	buff := make([]byte, fi.Size())
	//
	f, e := os.OpenFile(file_name, os.O_RDONLY, 0644)
	if e != nil {
		log.Fatalln("Open Error::", e.Error())
	}
	//
	n, e := f.Read(buff)
	if e != nil && e != io.EOF {
		log.Fatalln("Read Error::", e.Error())
	}
	//
	fmt.Printf("File %s read, byte count %d\n", file_name, n)
	//
	if e = f.Close(); e != nil {
		log.Fatalln("Close Error::", e.Error())
	}
	//
	b, _ := pem.Decode(buff)
	fmt.Printf("Block Type is %s\n", b.Type)
	fmt.Printf("Block Size is %d(0x%x)\n", len(b.Bytes), len(b.Bytes))
	fmt.Printf("Block Type is %s\n", b.Type)
	fmt.Printf("Block Headers are\n%q\n", b.Headers)
	fmt.Printf("Block:\n%s\n", hex.Dump(b.Bytes))
	//
	c, e := x509.ParseCertificate(b.Bytes)
	if e != nil {
		log.Fatalln("Parse Error::", e.Error())
	}
	//
	fmt.Printf("Cert Version: %d(0x%04x)\n", c.Version, c.Version)
	fmt.Printf("Cert Serial Number: %d(0x%x)\n", c.SerialNumber, c.SerialNumber)
	fmt.Printf("Cert Basic Constraints Valid: %v\n", c.BasicConstraintsValid)
	fmt.Printf("Cert IsCA: %v\n", c.IsCA)
	fmt.Printf("Cert Subject:\n%s\n", hex.Dump(c.RawSubject))
	//
	fmt.Println("done.....")
}
