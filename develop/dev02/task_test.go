package main

import "testing"

// go test -v -count=1 .
func TestAdd(t *testing.T) {
	for _, v := range []struct {
		str string
		exp string
	}{
		{
			str: "a4bc2d5e",
			exp: "aaaabccddddde",
		},

		{
			str: "",
			exp: "",
		},

		{
			str: "45",
			exp: "",
		},

		{
			str: `qwe\45`,
			exp: "qwe44444",
		},
	} {
		res, _ := stringUnpacking(v.str)
		if res != v.exp {
			t.Logf("%q не равно %v\n", res, v.exp)
			t.Fail()
		}
	}

}
