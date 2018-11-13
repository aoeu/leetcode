package leetcode

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	testCases := []struct {
		jewels   string
		stones   string
		expected int
	}{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}

	for _, tc := range testCases {
		if n := numJewelsInStones(tc.jewels, tc.stones); n != tc.expected {
			t.Errorf("Expected %v jewels but received %v for %v", tc.expected, n, tc)
		}
	}
}