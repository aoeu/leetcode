package leetcode

func toLowerCase(str string) string {
	out := make([]rune, len(str))
	for i, r := range str {
		if r >= 'A' && r <= 'Z' {
			r += 'a' - 'A'
		}
		out[i] = r
	}
	return string(out)
}

/*

https://leetcode.com/problems/to-lower-case/

Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.



Example 1:

Input: "Hello"
Output: "hello"
Example 2:

Input: "here"
Output: "here"
Example 3:

Input: "LOVELY"
Output: "lovely"

*/