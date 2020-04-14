/*
 * Maximum Subarray
 *
 * Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum
 * and return its sum.
 *
 * Example:
 *
 * Input: [-2,1,-3,4,-1,2,1,-5,4],
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 * Follow up:[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
 *
 * If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which
 * is more subtle.
 */

// NOTE: not sure if they meant to imply the divide and conquer approach is faster, but it's O(n log n). O(n) is the
// optimal solution for this problem.

package challenge

import "math"

func maxSubArray(nums []int) int {
	bestSum := math.MinInt32
	currSum := 0

	for _, number := range nums {
		currSum = max(number, currSum+number)
		bestSum = max(currSum, bestSum)
	}
	return bestSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
