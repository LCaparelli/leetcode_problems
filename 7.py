"""
7. Reverse Integer

Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321

Example 2:

Input: -123
Output: -321

Example 3:

Input: 120
Output: 21

Note: Assume we are dealing with an environment which could only store integers within the 32-bit signed integer
range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed
integer overflows.
"""


class Solution:
    def reverse(self, x: int) -> int:
        str_x = str(x)
        rev_x = "-" + str_x[:0:-1] if str_x[0] == "-" else str_x[::-1]
        int_rev = int(rev_x)
        return int_rev if -2147483648 <= int_rev <= 2147483647 else 0


sol = Solution()
print(sol.reverse(-120))
print(sol.reverse(-1201))
print(sol.reverse(120))
print(sol.reverse(1201))
print(sol.reverse(1534236469))
