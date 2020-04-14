/*
 * Single Number
 *
 * Given a non-empty array of integers, every element appears twice except for one. Find that single one.
 * Note:
 * Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
 *
 * Example 1:
 *
 * Input: [2,2,1]
 * Output: 1
 *
 * Example 2:
 *
 * Input: [4,1,2,1,2]
 * Output: 4
 */

package challenge

// First attempt
func singleNumber(nums []int) int {
	appearedOnce := make(map[int]bool)

	for _, number := range nums {
		if appearedOnce[number] {
			delete(appearedOnce, number)
		} else {
			appearedOnce[number] = true
		}
	}

	var answer int
	for number := range appearedOnce {
		answer = number
	}
	return answer
}

// Best solution
func bestSingleNumber(nums []int) int {
	a := 0
	for _, number := range nums {
		a ^= number
	}
	return a
}
