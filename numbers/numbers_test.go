/*
*/
package numbers
//
// Could import multiple packages here.  And also 'rename' the package for local
// use.
//
import (
	"testing"
)
//
// Declcare a type.  It is a struct.
//
type doubleTest struct {
	in, out int
}
//
// A literal declaration.  An array of types.
//
var doubleTests = []doubleTest{
	doubleTest{1, 2},
	doubleTest{2, 4},
	doubleTest{-5, -10},
}
//
// A basic test function.  Make sure multiply works.
// Note: exported.
//
func TestDouble(t *testing.T) {
	for _, dt := range doubleTests {
		v := Double(dt.in)
		if v != dt.out {
			t.Errorf("Double(%d) = %d, want %d.", dt.in, v, dt.out)
		}
	}
}
