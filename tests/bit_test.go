package tests

import "github.com/fredericlemoine/bitset"
import "testing"
import "fmt"

func TestBitset(t *testing.T) {
	var l uint = 67
	b, b2, b3 := bitset.New(l), bitset.New(l), bitset.New(l)

	if b.Len() != l {
		t.Error(fmt.Sprintf("Bitset Length should be %d and is %d", l, b.Len()))
	}

	var i uint
	for i = 0; i < l; i++ {
		if i%2 == 0 {
			b.Set(i)
			b3.Set(i)
			if !b.Test(i) {
				t.Error("bit ", i, " of b should be set but is not")
			}
			if !b3.Test(i) {
				t.Error("bit ", i, " of b3 should be set but is not")
			}
		} else {
			b2.Set(i)
			if !b2.Test(i) {
				t.Error("bit ", i, " of b2 should be set but is not")
			}
		}
	}

	if !b.Complement().Equal(b2) {
		t.Error("b and b2 should be complement but are not")
	}
	if !b.EqualOrComplement(b2) {
		t.Error("b and b2 should be complement but are not")
	}
	if !b.Equal(b3) {
		t.Error("b and b3 should be equal but are not")
	}

	fmt.Println("b : " + b.DumpAsBits())
	fmt.Println("b2: " + b2.DumpAsBits())
	fmt.Println("b3: " + b3.DumpAsBits())
}
