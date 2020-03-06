/*
139. Word Break

Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be segmented into a space-separated sequence of one or more dictionary words.

Note:

    The same word in the dictionary may be reused multiple times in the segmentation.
    You may assume the dictionary does not contain duplicate words.

Example 1:

Input: s = "leetcode", wordDict = ["leet", "code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".

Example 2:

Input: s = "applepenapple", wordDict = ["apple", "pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
             Note that you are allowed to reuse a dictionary word.

Example 3:

Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
Output: false
*/

package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	dictSet := make(map[string]bool)
	for _, word := range wordDict {
		dictSet[word] = true
	}

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func main() {
	var expected bool

	expected = true
	fmt.Printf("Expected: %v\nResult: %v\n\n", expected, wordBreak("123456789", []string{"123", "456", "789"}))

	expected = true
	fmt.Printf("Expected: %v\nResult: %v\n\n", expected, wordBreak("applepenapple", []string{"apple", "pen"}))

	expected = false
	fmt.Printf("Expected: %v\nResult: %v\n\n", expected, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
