package leetcode

import "strings"

func numJewelsInStones(J string, S string) int {
	if J == "" || S == "" {
		return 0
	}
	// TODO(aoeu): Assert "all characters in J and S are letters," because the types here do not.
	jewels := make(map[string]struct{}, 0)
	for _, s := range strings.Split(J, "") {
		jewels[s] = struct{}{}
	}
	n := 0
	for _, s := range strings.Split(S, "") {
		if _, ok := jewels[s]; ok {
			n++
		}
	}
	return n
}

/*

https://leetcode.com/problems/jewels-and-stones/

771. Jewels and Stones
Easy
1038
187


You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.

The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".

Example 1:

Input: J = "aA", S = "aAAbbbb"
Output: 3
Example 2:

Input: J = "z", S = "ZZ"
Output: 0
Note:

S and J will consist of letters and have length at most 50.
The characters in J are distinct.

*/