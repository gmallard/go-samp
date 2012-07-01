/*
Convert ISO-8859-15 to UTF-8.
*/
package main

// This came from:
//
// http://groups.google.com/group/golang-nuts/browse_thread/thread/9e42a38620625a26
//
// as a way to convert ISO-8859-15 to UTF-8, .....
//
import "fmt"

// ISO-8859-15 to UTF-8
func ISO885915ToUTF8(c []byte) string {

	// http://unicode.org/Public/MAPPINGS/ISO8859/8859-15.TXT 03-Mar-2004 14:06

	u := make([]rune, len(c))
	for i := 0; i < len(u); i++ {
		r := rune(c[i])
		if r >= 0x80 {
			switch r {
			case 0xA4:
				r = 0x20AC // EURO SIGN
			case 0xA6:
				r = 0x0160 // LATIN CAPITAL LETTER S WITH CARON
			case 0xA8:
				r = 0x0161 // LATIN SMALL LETTER S WITH CARON
			case 0xB4:
				r = 0x017D // LATIN CAPITAL LETTER Z WITH CARON
			case 0xB8:
				r = 0x017E // LATIN SMALL LETTER Z WITH CARON
			case 0xBC:
				r = 0x0152 // LATIN CAPITAL LIGATURE OE
			case 0xBD:
				r = 0x0153 // LATIN SMALL LIGATURE OE
			case 0xBE:
				r = 0x0178 // LATIN CAPITAL LETTER Y WITH DIAERESIS
			}
		}
		u[i] = r
	}
	return string(u)

}

func main() {
	c := []byte{'0', 'A', 'a', 0xA4, 0xA6, 0xA8, 0xB4, 0xB8, 0xBC, 0xBD, 0xBE}
	s := ISO885915ToUTF8(c)
	fmt.Println(s)
	for _, cv := range s {
		fmt.Printf("%x ", cv)
	}
	fmt.Println()
	//
	fmt.Printf("%s\n", string(0xC0)) // Sanity check.  Is this really valid UTF-8?
}
