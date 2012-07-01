/*
Display PEM file information.
*/
package main

import (
  "crypto/x509"
  "encoding/pem"
  "fmt"
  "io"
  "log"
  "os"
)

var file_name = "../certs/server/ServerTJCA.crt"

func main() {
  fmt.Println("start.....")
  fi, e := os.Lstat(file_name)
  if e != nil {
    log.Fatalln("Lstat Error::", e)
  }
  fmt.Printf("File Size: %d\n", fi.Size())
  //
  buff := make([]byte, fi.Size())
  //
  f, e := os.OpenFile(file_name, os.O_RDONLY, 0644)
  if e != nil {
    log.Fatalln("Open Error::", e)
  }
  //
  n, e := f.Read(buff)
  if e != nil && e != io.EOF {
    log.Fatalln("Read Error::", e)
  }
  //
  fmt.Printf("File %s read, byte count %d\n", file_name, n)  
  //
  if e = f.Close(); e != nil {
    log.Fatalln("Close Error::", e)
  }
  //
  b, _ := pem.Decode(buff)
  fmt.Printf("Block Type is %s\n", b.Type)
  fmt.Printf("Block Size is %d\n", len(b.Bytes))
  //
  c, e := x509.ParseCertificate(b.Bytes)
  if e != nil {
    log.Fatalln("Parse Error::", e)
  }
  //
  fmt.Printf("Cert Version: %d\n", c.Version)
  fmt.Printf("Cert Serial Number: %d\n", c.SerialNumber)
  fmt.Printf("Cert Basic Constraints Valid: %v\n", c.BasicConstraintsValid)
  fmt.Printf("Cert IsCA: %v\n", c.IsCA)
  // fmt.Printf("Cert Subject: %s\n", string(c.RawSubject))
  //
  p := x509.NewCertPool()
  p.AddCert(c)
  fmt.Println("done.....")
}

