package main

import (
  "crypto/tls"
  "log"
)

var testConfig *tls.Config

var hap = "tjjackson:61612"

var client_cert = "../certs/client/ClientTJ.crt"

var client_key = "../certs/client/ClientTJ.key"


func dumpState(s tls.ConnectionState) {
  log.Printf("Handshake Complete: %v\n", s.HandshakeComplete)
  log.Printf("Negotiated Protocol: %s\n", s.NegotiatedProtocol)
  log.Printf("Negotiated Protocol Mutual: %v\n", s.NegotiatedProtocolIsMutual)
  log.Printf("Server Name: %s\n", s.ServerName)
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
  testConfig.Certificates = cert_list // Add certificate list

  // Connect
  conn, e := tls.Dial("tcp", hap, testConfig)
  // Check e
  if e != nil {
    log.Fatalln("Dial Error::", e)
  }
  //
  log.Println("point 01")
  s := conn.ConnectionState()
  dumpState(s)
  if !s.HandshakeComplete {
    e = conn.Handshake()
    if e != nil {
      log.Fatalln("Handshake Error::", e)
    }
    log.Println("point 02")
    dumpState(conn.ConnectionState())
  }
  log.Println("point 03")
  //
  e = conn.Close()
  if e != nil {
    log.Fatalln("Close Error::", e)
  }

  log.Println("done....")
}

