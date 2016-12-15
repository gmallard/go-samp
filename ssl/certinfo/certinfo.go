/*
Display some information about an SSL certificate.
*/
package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"log"
)

type certFileData struct {
	cert string
	key  string
}

var certFiles = []certFileData{
	{"../certs/client/client.crt", "../certs/client/client.key"},
}

func showX509Cert(c *x509.Certificate) {
	log.Println("RAW Length:", len(c.Raw))
	rds := hex.Dump(c.Raw)
	log.Printf("Rawdata:\n%s\n", rds)
}

func showTLSCert(c tls.Certificate) {
	log.Println("===============================================================")
	log.Println("Chain Length:", len(c.Certificate))
	for n, cx := range c.Certificate {
		log.Println("Cert Element: ", n+1, "Length:", len(cx))
		log.Printf("%s\n", hex.Dump(cx))
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
		log.Println("Cert number:", n+1)
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
