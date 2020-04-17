/*
 * Perform String Shifts
 *
 * You are given a string s containing lowercase English letters, and a matrix shift, where shift[i] = [direction, amount]:
 *
 * direction can be 0 (for left shift) or 1 (for right shift).
 * amount is the amount by which string s is to be shifted.
 * A left shift by 1 means remove the first character of s and append it to the end.
 * Similarly, a right shift by 1 means remove the last character of s and add it to the beginning.
 *
 * Return the final string after all operations.
 *
 *
 *
 * Example 1:
 *
 * Input: s = "abc", shift = [[0,1],[1,2]]
 * Output: "cab"
 * Explanation:
 * [0,1] means shift to left by 1. "abc" -> "bca"
 * [1,2] means shift to right by 2. "bca" -> "cab"
 *
 * Example 2:
 *
 * Input: s = "abcdefg", shift = [[1,1],[1,1],[0,2],[1,3]]
 * Output: "efgabcd"
 * Explanation:
 * [1,1] means shift to right by 1. "abcdefg" -> "gabcdef"
 * [1,1] means shift to right by 1. "gabcdef" -> "fgabcde"
 * [0,2] means shift to left by 2. "fgabcde" -> "abcdefg"
 * [1,3] means shift to right by 3. "abcdefg" -> "efgabcd"
 *
 *
 *
 * Constraints:
 *
 * 1 <= s.length <= 100
 * s only contains lower case English letters.
 * 1 <= shift.length <= 100
 * shift[i].length == 2
 * 0 <= shift[i][0] <= 1
 * 0 <= shift[i][1] <= 100
 */

package challenge

import (
	"bytes"
)

const (
	LEFT = 0
)

func stringShift(s string, shifts [][]int) string {
	b := new(bytes.Buffer)
	shift := totalShift(shifts) % len(s)

	if shift == 0 {
		return s
	}

	for i := 0; i < len(s); i++ {
		j := (i + abs(len(s)-shift)) % len(s)
		b.Write([]byte{s[j]})
	}
	return b.String()
}

func totalShift(shifts [][]int) int {
	var total int
	for _, shift := range shifts {
		if shift[0] == LEFT {
			total -= shift[1]
		} else {
			total += shift[1]
		}
	}
	return total
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
