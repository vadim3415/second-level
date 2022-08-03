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
func TestAdd(t *testing.T) {
	for _, v := range []struct {
		exp  []string
		k    int
		r    bool
		u    bool
		n    bool
		file string
	}{
		{
			exp:  []string{"виноград 3 вино", "брусчатка 2 бардюр", "аквариум 1 вода"},
			k:    2,
			r:    true,
			u:    true,
			n:    false,
			file: "note.txt",
		},

		{
			exp: []string{"аквариум 1 вода", "аквариум 1 вода", "брусчатка 2 бардюр",
				"брусчатка 2 бардюр", "виноград 3 вино", "виноград 3 вино",
			},
			k:    2,
			r:    false,
			u:    false,
			n:    false,
			file: "note.txt",
		},
	} {
		res := sorted(v.file, v.k, v.r, v.u, v.n)
		if !Equal(v.exp, res) {
			t.Logf("%q не равно %q\n", res, v.exp)
			t.Fail()
		}
	}

}
