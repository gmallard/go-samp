/*
Reading XML.  Demonstrate load to a slice.
*/
package main

import "fmt"
import "encoding/xml"

var data string = `
<table>
    <transport>
        <uri>tcp://0.0.0.0:1234</uri>
    </transport>
    <transport>
        <uri>tls://127.0.0.1:7654?a=b</uri>
    </transport>
</table>
`

type Transport struct {
	Uri string `xml:"uri"`
}

type Table struct {
	XMLName    xml.Name    `xml:"table"`
	Transports []Transport `xml:"transport"`
}

func main() {
	var table Table
	err := xml.Unmarshal([]byte(data), &table)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(table)
	fmt.Println("1", table.Transports[0].Uri)
	fmt.Println("2", table.Transports[1].Uri)
}
