package leetcode

import "testing"

func TestToLowerCase(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{"Hello", "hello"},
		{"here", "here"},
		{"LOVELY", "lovely"},
	}
	for _, tc := range testCases {
		if s := toLowerCase(tc.in); s != tc.out {
			t.Errorf("Expected '%v' but received '%v' for input '%v'", tc.out, s, tc.in)
		}
	}
}