package leetcode

import "testing"

func TestJudgeCircle(t *testing.T) {
	testCases := []struct {
		in  string
		out bool
	}{
		{"UD", true},
		{"LL", false},
	}
	for _, tc := range testCases {
		if out := judgeCircle(tc.in); out != tc.out {
			t.Errorf("Expected %v but received %v for input '%v'", tc.out, out, tc.in)
		}
	}
}