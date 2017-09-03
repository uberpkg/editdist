package editdist

import (
	"fmt"
	"testing"
)

func TestMinimum(t *testing.T) {
	tests := []struct {
		a, b, c int
		exp     int
	}{
		{a: 0, b: 0, c: 0, exp: 0},
		{a: 1, b: 5, c: 7, exp: 1},
		{a: 5, b: 1, c: 7, exp: 1},
		{a: 6, b: 8, c: 3, exp: 3},
		{a: 6, b: 5, c: 4, exp: 4},
		{a: 6, b: 5, c: 4, exp: 4},
		{a: 6, b: 6, c: 4, exp: 4},
		{a: 6, b: 5, c: 6, exp: 5},
		{a: 3, b: 5, c: 5, exp: 3},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d %d %d", test.a, test.b, test.c), func(t *testing.T) {
			res := minimum(test.a, test.b, test.c)
			if res != test.exp {
				t.Errorf("minimum(%d, %d, %d) = %d, expected %d",
					test.a, test.b, test.c, res, test.exp)
			}
		})
	}
}

func TestLevenshtein(t *testing.T) {
	tests := []struct {
		a, b string
		exp  int
	}{
		{a: "kitten", b: "sitting", exp: 3},
		{a: "kitten", b: "", exp: 6},
		{a: "", b: "sitting", exp: 7},
		{a: "cat", b: "cat", exp: 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s %s", test.a, test.b), func(t *testing.T) {
			res := Levenshtein(test.a, test.b)
			if res != test.exp {
				t.Errorf("Levenshtein(%s, %s) = %d, expected %d",
					test.a, test.b, res, test.exp)
			}
		})
	}
}

func TestWagnerFischer(t *testing.T) {
	tests := []struct {
		a, b string
		exp  int
	}{
		{a: "kitten", b: "sitting", exp: 3},
		{a: "kitten", b: "", exp: 6},
		{a: "", b: "sitting", exp: 7},
		{a: "cat", b: "cat", exp: 0},
		{a: "Sunday", b: "Saturday", exp: 3},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s %s", test.a, test.b), func(t *testing.T) {
			res := WagnerFischer(test.a, test.b)
			if res != test.exp {
				t.Errorf("Levenshtein(%s, %s) = %d, expected %d",
					test.a, test.b, res, test.exp)
			}
		})
	}
}

func TestHjelmqvist(t *testing.T) {
	tests := []struct {
		a, b string
		exp  int
	}{
		{a: "kitten", b: "sitting", exp: 3},
		{a: "kitten", b: "", exp: 6},
		{a: "", b: "sitting", exp: 7},
		{a: "cat", b: "cat", exp: 0},
		{a: "Sunday", b: "Saturday", exp: 3},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s %s", test.a, test.b), func(t *testing.T) {
			res := Hjelmqvist(test.a, test.b)
			if res != test.exp {
				t.Errorf("Levenshtein(%s, %s) = %d, expected %d",
					test.a, test.b, res, test.exp)
			}
		})
	}
}
