/*
Show a more interesting use in interface{}.
A function is passed two very different type of interfacesm and figures
ut which to use.
*/
package main

import (
	"log"
	"net"
	"reflect"
)

type Figer interface {
	Configure()
}

type Configuration struct {
	Figer
	NetConn net.Conn
	Ival    int
}

func (c Configuration) Configure() {
	log.Println("Configure runs:", c.Ival)
}

var (
	mc  net.Conn
	cfg Configuration
	err error
)

func Taker(as string, ai interface{}) {
	log.Println("====================================")
	log.Println("Taker runs", as)
	log.Println("ai is a:", reflect.TypeOf(ai))
	isnc, iscfg := false, false
	switch mytype := ai.(type) {
	case net.Conn:
		log.Println("found net.Conn", reflect.TypeOf(mytype))
		isnc = true
	case Configuration:
		log.Println("found Configuration", reflect.TypeOf(mytype))
		iscfg = true
	default:
		panic(reflect.TypeOf(ai))
	}
	log.Println(isnc, iscfg)
	log.Println("====================================")
	if isnc {
		tc := ai.(*net.TCPConn)
		log.Println(tc.RemoteAddr())
	} else {
		cfg := ai.(Configuration)
		cfg.Configure()
	}
}

func main() {
	log.Println("start")
	cfg = Configuration{Ival: 42}
	cfg.Configure()
	//
	mc, err = net.Dial("tcp", "localhost:5900")
	if err != nil {
		panic(err)
	}
	Taker("net.Conn", mc)
	Taker("Configuration", cfg)
	log.Println("end")
}
