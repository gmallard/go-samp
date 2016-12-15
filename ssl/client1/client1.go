/*
SSL Use Case 1.
*/
package main

import (
	"crypto/tls"
	"fmt"
	"log"
)

var (
	testConfig *tls.Config
	hapnocl    = "localhost:61611" // ActiveMQ (needClientAuth="false")
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
// SSL Use Case 1 - server does *not* authenticate client, client does *not* authenticate server
//
// Subcase 1.A - Server configuration does *not* require client authentication
//
// - Expect connection success
//
// Subcase 1.B - Server configuration *does* require client authentication
//
// - Expect connection failure (server must be sent a valid client certificate)
//

func main() {
	fmt.Println("start......")
	testConfig = new(tls.Config)
	testConfig.InsecureSkipVerify = true         // Do *not* check the server's certificate
	c, e := tls.Dial("tcp", hapnocl, testConfig) // Server does not require a cert
	// from us.
	if e != nil {
		log.Fatalln("Dial Error::", e.Error())
	}
	//
	log.Println("Dial complete")
	log.Println("Connection State Reference:")
	s := c.ConnectionState()
	dumpState(s)
	if !s.HandshakeComplete {
		e = c.Handshake()
		if e != nil {
			log.Fatalln("Handskake Not Complete! Handshake Error:", e.Error())
		}
	}
	log.Println("Handshake Complete OK")
	//
	e = c.Close()
	if e != nil {
		log.Fatalln("Close Error::", e)
	}
	// Check e
	fmt.Println("done......")
}
