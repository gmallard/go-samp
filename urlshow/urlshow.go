/*
Demo handling URLs .....
*/
package main

import (
	"fmt"
	"net/url"
)

/*
References:

	https://tools.ietf.org/html/rfc1738

	https://www.w3schools.com/tags/ref_urlencode.asp

*/

func dumpURL(d string) {
	fmt.Println()
	fmt.Println("==================")
	fmt.Println("Data:", d)

	u, e := url.Parse(d)
	if e != nil {
		panic(e)
	}

	fmt.Println("Raw:", u)

	fmt.Println("IsAbs:", u.IsAbs())

	fmt.Println("RequestURI", u.RequestURI())

	if u.Scheme != "" {
		fmt.Println("Scheme:", u.Scheme)
	}

	if u.Opaque != "" {
		fmt.Println("Qpaque:", u.Opaque)
	}

	if u.User != nil {
		fmt.Println("User:", u.User)
	}

	if u.Host != "" {
		fmt.Println("Host:", u.Host)
		fmt.Println("\tHostname", u.Hostname())
		fmt.Println("\tPort", u.Port())
	}

	if u.Path != "" {
		fmt.Println("Path:", u.Path)
		fmt.Println("\tEscapedPath:", u.EscapedPath())
	}

	if u.RawQuery != "" {
		fmt.Println("rawquery:", u.RawQuery)
		v := u.Query()
		for ok, ov := range v {
			fmt.Print("key ", ok, " vals ")
			for _, nv := range ov {
				fmt.Print(" ", nv)
			}
			fmt.Println()
		}
	}

	if u.Fragment != "" {
		fmt.Println("Fragment:", u.Fragment)
	}

}

func main() {
	tests := []string{"localhost.com",
		"localhost.com/",
		"tcp://localhost.com:9191",
		"tcp://localhost.com:9191/",
		"tls://example.com:1337/?a=1",
		"tcp6://example6.com:1338/?a=1&b=2",
		"stomp://somewhere.org:11613",
		"tls://example.com:1337/?a=1#fragment1",
		"tls://example.com:1337/?a=1%23fragment1",
		"tls://example.com:1337/?a=1#fragment1&b=2#fragment2",
		"tls://example.com:1111/?a=1&b=2#fragment1",
		"tls://example.com:1111/?a=1&b=2%23fragment1",
		"tls://example.com:1111/?a=1&b=2#fragment1#fragment2",
		"tls://example.com:1111/?a=1&b=2%23fragment1%23fragment2",
		"tcp6://user6@example.com:2222",
		"tls://examplez.com:3333/?a=1+2+3",
		"tls://examplez.com:4321/pa/pb/?a=1+2+3",
		"tcp+ssl://examples.com:9875/?a=1+2+3",
		"/examples.com:3333",
		"examples2.com:2222",
	}
	//
	// Some of the results here are a little surprising.
	// Especially the tests wih:
	// a) fragment(s)
	// b) no scheme
	// Inspect them carefully.
	//
	for _, v := range tests {
		dumpURL(v)
	}
}
