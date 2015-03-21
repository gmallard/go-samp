/*
Demo handling URLs .....
*/
package main

import (
	"fmt"
	"net/url"
)

func dumpURL(d string) {
	fmt.Println("==================")
	fmt.Println("Data:", d)

	u, e := url.Parse(d)
	if e != nil {
		panic(e)
	}

	fmt.Println("Raw:", u)

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
	}

	if u.Path != "" {
		fmt.Println("Path:", u.Path)
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
		"tls://example.com:1337/?a=1#fragment1&b=2#fragment2", // surprise
		"tls://example.com:1111/?a=1&b=2#fragment1",
		"tcp6://user6@example.com:2222",
		"tls://examplez.com:3333/?a=1+2+3",
	}
	for _, v := range tests {
		dumpURL(v)
	}
}
