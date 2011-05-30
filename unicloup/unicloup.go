// A conundrum in UTF-8 upper/lower case .........

package main

import (
	"fmt" //
  "strings"
)
//
// An interesting exercise with UTF-8.  Triggered by this discussion:
// http://stackoverflow.com/questions/6162484/why-does-modern-perl-avoid-utf-8-by-default/#6163129
//
// I also experimented with this using Ruby 1.9.3, the current tip of trunk.
// Ruby seems to handle this even less gracefully than go does, although
// the Ruby documentation clearly says: String#upcase and String#downcase
// operate on the 'ASCII range' only. Thus, the code:
// "σ".upcase
// does absolutely nothing.
//
// There is an interesting presentation on Unicode work by the PHP development
// team here:
// http://www.slideshare.net/andreizm/the-good-the-bad-and-the-ugly-what-happened-to-unicode-and-php-6
// One of the interesting slides is number 53 where an example showing "Σ"
// being 'lower case' converted to either "σ" or "ς".  The conclusions of
// the slide deck are however, discouraging.
// 
//
func main() {
	fmt.Println("Start...")

  // Lower case sigma
  lowsig := "σ"
  fmt.Printf("Low Sig: %s\n", lowsig)

  // Lower case C with cedilla
  lowceeced := "ς"
  fmt.Printf("Low Cee Ced: %s\n", lowceeced)

  // Upper case sigma
  hisig := "Σ"
  fmt.Printf("High Sig: %s\n", hisig)

  // Lower case sigma to upper
  fmt.Printf("Toupper Sig: %s\n", strings.ToUpper(lowsig))

  // Lower case C cedilla to upper
  fmt.Printf("Toupper Cee Ced: %s\n", strings.ToUpper(lowceeced))

  // Upper case sigma to lower
  fmt.Printf("Tolower Cee: %s\n", strings.ToLower(hisig))

	fmt.Println("End...")
}
