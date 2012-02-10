//
// The package name.  This name will become a .a file in $(GOROOT)/pkg.
//
package numbers
//
// Note function with name that is capitalized.  It _is_ exported.
//
func Double(i int) int {
	//
	// Not much to do here.  Just multiply by 2.
	//
	return i * 2
}
