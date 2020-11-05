/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

package main

import (
	"math"
)

// @lc code=start
func longestPalindrome(s string) string {
	childString := ""

	for i := 0.5; i < float64(len(s)); i += 0.5 {
		for j := 0.0; j < math.Min(float64(len(s))-i, i); j++ { // 2
			prev := int(math.Ceil(i - j))
			next := int(math.Floor(i + j))

			if s[prev] == s[next] {
				childString = s[prev : next+1]
			}
		}
	}

	return childString
}

// @lc code=end
