/*
 * Backspace String Compare
 *
 * Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.
 *
 * Note that after backspacing an empty text, the text will continue empty.
 *
 * Example 1:
 *
 * Input: S = "ab#c", T = "ad#c"
 * Output: true
 * Explanation: Both S and T become "ac".
 *
 * Example 2:
 *
 * Input: S = "ab##", T = "c#d#"
 * Output: true
 * Explanation: Both S and T become "".
 *
 * Example 3:
 *
 * Input: S = "a##c", T = "#a#c"
 * Output: true
 * Explanation: Both S and T become "c".
 *
 * Example 4:
 *
 * Input: S = "a#c", T = "b"
 * Output: false
 * Explanation: S becomes "c" while T becomes "b".
 *
 * Note:
 *
 * 1 <= S.length <= 200
 * 1 <= T.length <= 200
 * S and T only contain lowercase letters and '#' characters.
 *
 * Follow up:
 *
 * Can you solve it in O(N) time and O(1) space?
 */

package challenge

import (
	"bytes"
)

func backspaceCompare(s string, t string) bool {
	var bs, bt = new(bytes.Buffer), new(bytes.Buffer)
	compileExpr(bs, s)
	compileExpr(bt, t)

	s, t = bs.String(), bt.String()
	return bs.String() == bt.String()
}

func compileExpr(b *bytes.Buffer, expr string) {
	expr = trimPrefix(expr)

	for _, char := range expr {
		if char == '#' {
			if b.Len() > 0 {
				b.Truncate(b.Len() - 1)
			}
		} else {
			b.Write([]byte{byte(char)})
		}
	}
}

// strings.trimPrefix can't be used with a pattern, so this needed to be implemented to trim '^#*'
func trimPrefix(s string) string {
	var i int
	for i = 0; s[i] == '#'; i++ {
		continue
	}
	return s[i:]
}
