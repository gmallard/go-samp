/*
SSL Use Case 3.
*/
package main

import (
	"crypto/tls"
	"log"
)

var (
	testConfig  *tls.Config
	hapcl       = "localhost:61612" // ActiveMQ (needClientAuth="true")
	client_cert = "../certs/client/client.crt"
	client_key  = "../certs/client/client.key"
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
// SSL Use Case 3 - server *does* authenticate client, client does *not* authenticate server
//
// Subcase 3.A - Message broker configuration does *not* require client authentication
//
// - Expect connection success
//
// Subcase 3.B - Message broker configuration *does* require client authentication
//
// - Expect connection success if the server can authenticate the client certificate
//
func main() {
	log.Println("start....")
	// Create a tls.Certificate structure from the clients cert and key
	cert, e := tls.LoadX509KeyPair(client_cert, client_key)
	if e != nil {
		log.Fatalln("Error: ", e)
	}
	// Create a certificate list, and add the client's certificate
	cert_list := make([]tls.Certificate, 1)
	cert_list[0] = cert
	// Create and initialize tls.Config
	testConfig = new(tls.Config)
	testConfig.InsecureSkipVerify = true // Do *not* check the server's certificate
	testConfig.Certificates = cert_list  // Add certificate list
	// Connect
	conn, e := tls.Dial("tcp", hapcl, testConfig)
	// Check e
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
	log.Println("Handshake Complete OK") //
	e = conn.Close()
	if e != nil {
		log.Fatalln("Close Error::", e)
	}
	log.Println("done....")
}
