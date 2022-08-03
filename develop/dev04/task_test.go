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
		exp map[string][]string
		str []string
	}{
		{
			exp: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"слиток": {"слиток", "листок", "столик"},
				"кино":   {"кино", "кони"},
			},
			str: []string{"кино", "кони", "Пятак", "пятка", "слиток",
				"листок", "столик", "листок", "Порт", "рог", "тяпка", "s",
			},
		},
	} {
		res := findAnagrams(v.str)
		if len(res) != len(v.exp) {
			t.Logf("%q не равно %q\n", res, v.exp)
			t.Fail()
			return
		}
		for key, val := range res {
			if !Equal(val, v.exp[key]) {
				t.Logf("%q не равно %q\n", res, v.exp)
				t.Fail()
			}
		}

	}

}
