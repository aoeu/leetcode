package leetcode

import (
	"reflect"
	"testing"
)

func SortArrayByParity(t *testing.T) {
	testCases := []struct {
		in  []int
		out []int
	}{

		{[]int{3, 1, 2, 4}, []int{2, 4, 3, 1}},
	}
	for _, tc := range testCases {
		if out := sortArrayByParity(tc.in); !reflect.DeepEqual(out, tc.out) {
			t.Errorf("Expected %v but received %v for input %v", tc.out, out, tc.in)
		}
	}
}