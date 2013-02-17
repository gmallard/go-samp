/*
A demonstration of gob encoding and decoding.
*/
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

type sample struct {
	Id      string
	Headers []string
	Body    []byte
}

// Persist the encoded data.
func writeFile(n string, d []byte) error {
	f, e := os.Create(n)
	if e != nil {
		return e
	}
	defer f.Close()
	nr, e := f.Write(d)
	if e != nil {
		return e
	}
	fmt.Println("WriteDataLen:", len(d), "Written:", nr)
	return nil
}

// Load the encoded data.
func readFile(n string) ([]byte, error) {
	fi, e := os.Lstat(n)
	if e != nil {
		return nil, e
	}
	f, e := os.Open(n)
	if e != nil {
		return nil, e
	}
	defer f.Close()
	sz := int(fi.Size())
	b := make([]byte, sz)
	nr, e := f.Read(b)
	if e != nil {
		return nil, e
	}
	fmt.Println("ReadDataLen:", len(b), "Read:", nr)
	return b, nil
}

func main() {
	// Sample data
	s := sample{"NAMEA", []string{"str1", "str2", "str3", "str4"},
		[]byte{0xa, 0xb, 0xc}}
	fmt.Printf("%+v\n", s)

	//  Encode.
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(s)
	if err != nil {
		fmt.Println("enc.Encode failed:", err)
		return
	}

	// Persist and Reload.
	err = writeFile("./tempdata.dat", buf.Bytes())
	if err != nil {
		fmt.Println("writeFile failed:", err)
		return
	}
	b, err := readFile("./tempdata.dat")
	if err != nil {
		fmt.Println("readFile failed:", err)
		return
	}

	// Decode.
	buf = bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	var d sample
	err = dec.Decode(&d)
	if err != nil {
		fmt.Println("dec.Decode failed:", err)
		return
	}
	fmt.Printf("%+v\n", d)
}
