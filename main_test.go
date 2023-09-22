package main

import (
	"testing"
)

func TestRotationalCipher(t *testing.T) {
	testCases := []struct {
		name     string
		value    string
		expected string
	}{
		{name: "all-lower-case", value: "abc", expected: "bcd"},
		{name: "all-lower-case-with-comma", value: "a,b,c", expected: "b,c,d"},
		{name: "all-upper-case", value: "ABCD", expected: "BCDE"},
		{name: "all-upper-case-with-comma", value: "A,B,C,D", expected: "B,C,D,E"},
		{name: "mix-upper-lower", value: "AbCd", expected: "BcDe"},
		{name: "mix-upper-lower-with-comman", value: "A,bCd,", expected: "B,cDe,"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := RotationalCipher(tc.value, 1)
			if got != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, got)
			}
		})
	}
}

func BenchmarkRotationalCipher(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RotationalCipher("abc", 10)
	}
}
