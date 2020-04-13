/*
 * Happy Number
 *
 * Write an algorithm to determine if a number n is "happy".
 *
 * A happy number is a number defined by the following process: Starting with any positive integer, replace the number
 * by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it
 * loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy
 * numbers.
 *
 * Return True if n is a happy number, and False if not.
 *
 * Example:
 *
 * Input: 19
 * Output: true
 * Explanation:
 * 12 + 92 = 82
 * 82 + 22 = 68
 * 62 + 82 = 100
 * 12 + 02 + 02 = 1
 */

package challenge

import "fmt"

func isHappy(n int) bool {
	fmt.Printf("Input: %d\n", n)
	hasAppeared := make(map[int]bool)
	digs := digits(n)

	// emulating a do-while
	var squaredSum int
	for _, digit := range digs {
		squaredSum += digit * digit
	}
	for squaredSum != 1 {
		squaredSum = 0
		for _, digit := range digs {
			squaredSum += digit * digit
		}

		// a second appearance means there's a cycle
		if hasAppeared[squaredSum] {
			return false
		}
		hasAppeared[squaredSum] = true
		digs = digits(squaredSum)
	}
	return true
}

func digits(n int) []int {
	currNum := n
	var ans []int

	for currNum > 0 {
		ans = append(ans, currNum%10)
		currNum = currNum / 10
	}
	return ans
}
