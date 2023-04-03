package main

import (
	"testing"
)

var tests = []struct {
	num  int
	want bool
}{
	{num: 121, want: true},
	{num: -121, want: false},
	{num: -100, want: false},
}

func TestIsPalindrom(t *testing.T) {
	for _, test := range tests {
		got := isPalindrome(test.num)
		if got != test.want {
			t.Errorf("isPalindrom %d should be %t, got %t", test.num, test.want, got)
		}
	}
}
