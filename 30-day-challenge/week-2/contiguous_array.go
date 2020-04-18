/*
 * Contiguous Array
 *
 * Given a binary array, find the maximum length of a contiguous subarray with equal number of 0 and 1.
 *
 * Example 1:
 *
 * Input: [0,1]
 * Output: 2
 * Explanation: [0, 1] is the longest contiguous subarray with equal number of 0 and 1.
 *
 * Example 2:
 *
 * Input: [0,1,0]
 * Output: 2
 * Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
 *
 * Note: The length of the given binary array will not exceed 50,000.
 */

package challenge

func findMaxLength(nums []int) int {
	var count, maxLength int
	occurrences := make(map[int]int)
	// whenever count is 0 we have encountered an subarray of length i+1. We must initialize this entry as -1 so we can
	// simply assign the subarray length-(-1) == length+1 to maxLength
	occurrences[0] = -1

	for i, num := range nums {
		if num == 1 {
			count++
		} else {
			count--
		}
		if start, ok := occurrences[count]; !ok {
			occurrences[count] = i
		} else {
			if i-start > maxLength {
				maxLength = i - start
			}
		}
	}
	return maxLength
}

func main() {
	findMaxLength([]int{0, 1, 0})
	findMaxLength([]int{0, 1})
}
