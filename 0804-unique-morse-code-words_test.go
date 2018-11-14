package leetcode

import "testing"

func TestUniqueMorseRepresentations(t *testing.T) {
	testCases := []struct {
		in  []string
		out int
	}{
		{[]string{"gin", "zen", "gig", "msg"}, 2},
	}
	for _, tc := range testCases {
		if n := uniqueMorseRepresentations(tc.in); n != tc.out {
			t.Errorf("Expected a count of %v but received %v for input %v", tc.out, n, tc.in)
		}
	}
}