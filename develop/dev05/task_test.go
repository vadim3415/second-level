package main

import (
	"testing"
)

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// go test -v -count=1 .
func TestAfter(t *testing.T) {
	for _, v := range []struct {
		exp    []string
		file   string
		subStr string
		A      int
		B      int
		C      int
		c      bool
		i      bool
		v      bool
		F      bool
		n      bool
	}{
		{
			exp:    []string{"4: Autumn wyellow leaves cool", "5: Autumn Wyellow leaves cool2"},
			file:   "note.txt",
			subStr: "Wyellow",
			A:      1,
			B:      0,
			C:      0,
			c:      false,
			i:      true,
			v:      false,
			F:      false,
			n:      true,
		},
		{
			exp:    []string{"3: Summer ccolorful blossom hot", "4: Autumn wyellow leaves cool"},
			file:   "note.txt",
			subStr: "Wyellow",
			A:      0,
			B:      1,
			C:      0,
			c:      false,
			i:      true,
			v:      false,
			F:      false,
			n:      true,
		},
		{
			exp:    []string{"3: Summer ccolorful blossom hot", "4: Autumn wyellow leaves cool", "5: Autumn Wyellow leaves cool2"},
			file:   "note.txt",
			subStr: "Wyellow",
			A:      0,
			B:      0,
			C:      1,
			c:      false,
			i:      true,
			v:      false,
			F:      false,
			n:      true,
		},
		{
			exp:    []string{"count: 2"},
			file:   "note.txt",
			subStr: "wyellow",
			A:      0,
			B:      0,
			C:      0,
			c:      true,
			i:      false,
			v:      false,
			F:      false,
			n:      false,
		},
	} {
		res := grep(v.file, v.subStr, v.A, v.B, v.C, v.c, v.i, v.v, v.F, v.n)
		if len(res) != len(v.exp) {
			t.Logf("%q не равно %q\n", res, v.exp)
			t.Fail()
			return
		}

		if !Equal(res, v.exp) {
			t.Logf("%q не равно %q\n", res, v.exp)
			t.Fail()
		}
	}

}
