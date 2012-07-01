/*
Display some information about an SSL certificate.
*/
package main

import (
  "crypto/tls"
  "crypto/x509"
  "log"
)

type certFileData struct {
  cert  string
  key   string  
}

var certFiles = []certFileData{
    {"../certs/client/ClientTJ.crt", "../certs/client/ClientTJ.key"},
  }

func showX509Cert(c *x509.Certificate) {
  log.Println("RAW Length:", len(c.Raw))
}
 
func showTLSCert(c  tls.Certificate) {
  log.Println("Chain Length:", len(c.Certificate))
  for n, cx := range c.Certificate {
    log.Println("Element: ", n + 1, "Length:", len(cx))
    // log.Printf("Dump: %q\n", cx)
  }
  log.Println("OCSPStaple Length:", len(c.OCSPStaple))
  if c.Leaf == nil {
    log.Println("Leaf pointer is nil")
  } else {
    showX509Cert(c.Leaf)
  }
}

func main() {
  for n, cf := range certFiles {
    log.Println("Cert number:", n + 1)
    log.Println("Cert File:", cf.cert)
    log.Println("Key File:", cf.key)
    //
    c, e := tls.LoadX509KeyPair(cf.cert, cf.key)
    if e != nil {
      log.Fatalln("Error: ", e)
    }
    showTLSCert(c)
  }
}

