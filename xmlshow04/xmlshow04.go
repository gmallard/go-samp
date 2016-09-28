/*
	Reading XML.  Demonstrate recursive elements.
*/
package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

var data string = `
<circle>
	<xcoord>0</xcoord>
	<ycoord>0</ycoord>
	<radius>4</radius>
	<circle>
		<comment>Inside One</comment>
		<xcoord>2</xcoord>
		<ycoord>0</ycoord>
		<radius>2</radius>
		<circle>
			<comment>Inside Two</comment>
			<xcoord>1</xcoord>
			<ycoord>0</ycoord>
			<radius>1</radius>
		</circle>
	</circle>
</circle>
`

type Circle struct {
	XMLName xml.Name `xml:"circle"`
	Comment string   `xml:"comment"`
	Xcoord  int      `xml:"xcoord"`
	Ycoord  int      `xml:"ycoord"`
	Radius  int      `xml:"radius"`
	Circle  *Circle
}

func main() {
	var circle Circle
	err := xml.Unmarshal([]byte(data), &circle)
	if err != nil {
		log.Fatalf("Unmarshal ERROR: [%v]\n", err)
	}
	// fmt.Printf("Circle Data: [%q]\n", circle)
	fmt.Println("=============================")
	dumpData(&circle, "The Outside Circle")
}

func dumpData(c *Circle, id string) {
	fmt.Printf("A Circle (%s):\n", id)
	if c.Comment != "" {
		fmt.Printf("\tComment: %s\n", c.Comment)
	}
	fmt.Printf("\tX coordinate: %d\n", c.Xcoord)
	fmt.Printf("\tY coordinate: %d\n", c.Ycoord)
	fmt.Printf("\tRadius: %d\n", c.Radius)
	//
	if c.Circle != nil {
		dumpData(c.Circle, "Next Inner Circle")
	}
}
