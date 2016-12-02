/*
   Demonstrate a map[string]map[string]string:
   a map of a string to map of a string to  a string.
*/

package main

import "fmt"

func main() {
	type mss map[string]string
	mosa := mss{"asa": "astra", "asb": "astrb"}
	fmt.Println("mosa", mosa)
	mosb := mss{"bsa": "bstra", "bsb": "bstrb"}
	fmt.Println("mosb", mosb)
	//
	type msmss map[string]mss
	mossab := msmss{"mosa": mosa, "mosb": mosb}
	fmt.Println("mossab", mossab)
	//
	fmt.Println("------- mosa ----------")
	for k, v := range mosa {
		fmt.Printf("key=%s, value=%s\n", k, v)
	}
	fmt.Println("------- mosb ----------")
	for k, v := range mosb {
		fmt.Printf("key=%s, value=%s\n", k, v)
	}
	fmt.Println("------ mossab ---------")
	for ok, ov := range mossab {
		for ik, iv := range ov {
			fmt.Printf("ok=%s, ik=%s, iv=%s\n", ok, ik, iv)
		}
	}
}
