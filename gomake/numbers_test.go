package numbers

import (
	"testing"
)

type doubleTest struct {
	in, out int
}

var doubleTests = []doubleTest{
	doubleTest{1, 2},
	doubleTest{2, 4},
	doubleTest{-5, -10},
}

func TestDouble(t *testing.T) {
	for _, dt := range doubleTests {
		v := Double(dt.in)
		if v != dt.out {
			t.Errorf("Double(%d) = %d, want %d.", dt.in, v, dt.out)
		}
	}
}

