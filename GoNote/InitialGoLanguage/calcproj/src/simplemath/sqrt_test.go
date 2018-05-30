package simplemath

import "testing"

func TestSqrt1(t *testing.T) {
	r := Sqrt(16)

	if r != 4 {
		t.Errorf("Sqrt(16) faild, Got %d, expected 4", r)

	}

}
