/*
SSL Use Case 1.
*/
package main

import (
  "crypto/tls"
  "fmt"
  "log"
)

var testConfig *tls.Config

var hap = "tjjackson:61612"

func dumpState(s tls.ConnectionState) {
  log.Printf("Handshake Complete: %v\n", s.HandshakeComplete)
  log.Printf("Negotiaged Protocol: %s\n", s.NegotiatedProtocol)
  log.Printf("Negotiaged Protocol Mutual: %v\n", s.NegotiatedProtocolIsMutual)
  log.Printf("Server Name: %s\n", s.ServerName)
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
	testConfig.InsecureSkipVerify = true  // Do *not* check the server's certificate

  c, e := tls.Dial("tcp", hap, testConfig)

  if e != nil {
    log.Fatalln("Dial Error::", e)
  }
  //
  log.Println("point 01")
  s := c.ConnectionState()
  dumpState(s)

  if !s.HandshakeComplete {
    e = c.Handshake()
    if e != nil {
      log.Fatalln("Handshake Error::", e)
    }
    log.Println("point 02")
    dumpState(c.ConnectionState())
  }

  log.Println("point 03")
  //
  e = c.Close()
  if e != nil {
    log.Fatalln("Close Error::", e)
  }
  // Check e
  fmt.Println("done......")
}

