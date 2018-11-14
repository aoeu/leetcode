package leetcode

func sortArrayByParity(A []int) []int {
	evens, odds := make([]int, 0), make([]int, 0)
	for _, a := range A {
		if a%2 == 0 {
			evens = append(evens, a)
		} else {
			odds = append(odds, a)
		}
	}
	return append(evens, odds...)
}

/*

905. Sort Array By Parity
Easy
237
27


Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.

You may return any answer array that satisfies this condition.



Example 1:

Input: [3,1,2,4]
Output: [2,4,3,1]
The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.


Note:

1 <= A.length <= 5000
0 <= A[i] <= 5000

*/