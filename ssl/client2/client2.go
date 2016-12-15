/*
SSL Use Case 2.
*/
package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	testConfig *tls.Config
	hapnocl    = "localhost:61611" // ActiveMQ (needClientAuth="false")
	file_name  = "../certs/ca/ca.crt"
)

func dumpState(s tls.ConnectionState) {
	log.Printf("Version: %d(0x%04x)\n", s.Version, s.Version)
	log.Printf("HandshakeComplete: %t\n", s.HandshakeComplete)
	log.Printf("DidResume: %t\n", s.DidResume)
	log.Printf("CipherSuite: %d(0x%04x)\n", s.CipherSuite, s.CipherSuite)
	log.Printf("NegotiagedProtocol: %s\n", s.NegotiatedProtocol)
	log.Printf("NegotiagedProtocolIsMutual: %t\n", s.NegotiatedProtocolIsMutual)
	log.Printf("Server Name: %s\n", s.ServerName)
	log.Printf("Length PeerCertificates: %d(0x%04x)\n", len(s.PeerCertificates),
		len(s.PeerCertificates))
	log.Printf("Length VerifiedChains: %d(0x%04x)\n", len(s.VerifiedChains),
		len(s.VerifiedChains))
	log.Printf("Length SignedCertificateTimestamps: %d(0x%04x)\n", len(s.SignedCertificateTimestamps),
		len(s.SignedCertificateTimestamps))
	log.Printf("Length OCSPResponse: %d(0x%04x)\n", len(s.OCSPResponse),
		len(s.OCSPResponse))
}

//
// SSL Use Case 2 - server does *not* authenticate client, client *does* authenticate server
//
// Subcase 2.A - Message broker configuration does *not* require client authentication
//
// - Expect connection success
//
// Subcase 2.B - Message broker configuration *does* require client authentication
//
// - Expect connection failure (broker must be sent a valid client certificate)
//
func main() {
	fmt.Println("start.....")
	// Load the cert for the server's CA
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
	// Create an X509 Certificate structure from loaded data
	b, _ := pem.Decode(buff)
	fmt.Printf("Block Type is %s\n", b.Type)
	fmt.Printf("Block Size is %d(0x%x)\n", len(b.Bytes), len(b.Bytes))
	//
	c, e := x509.ParseCertificate(b.Bytes)
	if e != nil {
		log.Fatalln("Parse Error::", e)
	}
	// Display some information about the server's CA certificate
	fmt.Printf("Cert Version: %d(0x%04x)\n", c.Version, c.Version)
	fmt.Printf("Cert Serial Number: %d(0x%x)\n", c.SerialNumber, c.SerialNumber)
	fmt.Printf("Cert Basic Constraints Valid: %v\n", c.BasicConstraintsValid)
	fmt.Printf("Cert IsCA: %v\n", c.IsCA)
	// Create a X509 Certificate pool, and add the server's CA certificate
	p := x509.NewCertPool()
	p.AddCert(c)
	// Set up TLS config
	testConfig = new(tls.Config)
	testConfig.InsecureSkipVerify = false // *Do* check the server's certificate
	testConfig.RootCAs = p                // Add the certificate pool
	// Connect
	conn, e := tls.Dial("tcp", hapnocl, testConfig)
	if e != nil {
		log.Fatalln("Dial Error::", e)
	}
	//
	log.Println("Dial complete")
	log.Println("Connection State Reference:")
	s := conn.ConnectionState()
	dumpState(s)
	if !s.HandshakeComplete {
		e = conn.Handshake()
		if e != nil {
			log.Fatalln("Handskake Not Complete! Handshake Error:", e.Error())
		}
	}
	log.Println("Handshake Complete OK")
	//
	e = conn.Close()
	if e != nil {
		log.Fatalln("Close Error::", e)
	}
	fmt.Println("done.....")
}
