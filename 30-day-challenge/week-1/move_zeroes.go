/*
 * Move Zeroes
 *
 * Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the
 * non-zero elements.
 *
 * Example:
 *
 * Input: [0,1,0,3,12]
 * Output: [1,3,12,0,0]
 *
 * Note:
 *
 * You must do this in-place without making a copy of the array.
 * Minimize the total number of operations.
 */

package challenge

func moveZeroes(nums []int) {
	var i int
	allZeroes := true

	// deals with arrays already ending with one or more zeroes
	for j := len(nums) - 1; j >= 0; j-- {
		if nums[j] != 0 {
			i = j
			allZeroes = false
			break
		}
	}
	if allZeroes {
		return
	}

	for ; i >= 0; i-- {
		if nums[i] == 0 {
			removeAndReplace(i, nums)
		}
	}
}

func removeAndReplace(i int, nums []int) {
	for ; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
	nums[len(nums)-1] = 0
}
