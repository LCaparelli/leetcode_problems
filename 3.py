"""
3. Longest Substring Without Repeating Characters

Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
"""


class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if len(s) < 2:
            return len(s)

        longest = 0
        i = 0
        j = 1

        used_chars = set()
        used_chars.add(s[i])

        while j < len(s):
            if s[j] in used_chars:
                used_chars.remove(s[i])
                i += 1
            else:
                used_chars.add(s[j])
                j += 1
                candidate_len = j - i
                longest = candidate_len if candidate_len > longest else longest

        return longest


solution = Solution()
print(solution.lengthOfLongestSubstring("ohomm"))
print(solution.lengthOfLongestSubstring("abcabcbb"))
print(solution.lengthOfLongestSubstring("bbbbb"))
print(solution.lengthOfLongestSubstring("pwwkew"))
print(solution.lengthOfLongestSubstring("au"))
