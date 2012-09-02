/*
Reading XML.  A second demonstration.
*/
package main

import (
	"encoding/xml"
	"fmt"
	"time"
)

// Superficially similar to the problem described at the link below:
//
// http://code.google.com/p/go/issues/detail?id=3688
//

type Bookings struct {
	From     time.Time `xml:"from,attr"`
	Location string    `xml:"location,attr"`
}

type Configs struct {
	XMLName  xml.Name `xml:"root"`
	Login    string   `xml:"login"`
	Password string   `xml:"password"`
	Bookings Bookings `xml:"getBookings"`
}

var buff []byte // File data

var Xdata = `
<root>
	<login>test</login>
	<password>password</password>
	<getBookings from="2011-01-01T12:00:00Z" location="30038"/>
</root>`

func showXml() {
	ctx := Configs{}
	err := xml.Unmarshal(buff, &ctx)
	if err != nil {
		fmt.Printf("unmarshal error: %v", err)
		return
	}
	fmt.Printf("Configs: %+v\n", ctx)
	//
	m, err := xml.MarshalIndent(ctx, "", "\t")
	if err != nil {
		fmt.Printf("marshal error: %v", err)
		return
	}
	fmt.Println()
	fmt.Println("same?", string(m) == string(buff))
	fmt.Println()
	fmt.Println("strdatam:", string(m))
	fmt.Println()
	fmt.Println("bytedatam:", m, "len:", len(m))
	fmt.Println()
	fmt.Println("strdatab:", string(buff))
	fmt.Println()
	fmt.Println("bytedatab:", buff, "len:", len(buff))
}

func main() {

	buff = []byte(Xdata)
	showXml()

}
