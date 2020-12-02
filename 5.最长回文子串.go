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
	if len(s) == 1 {
		return s
	}

	childString := s[0:1]

	for i := 0.5; i < float64(len(s)); i += 0.5 {
		for j := 0.0; i+j < float64(len(s))-1; j++ { // 2
			if i-j >= 0 {
				prev := int(math.Ceil(i - j))
				next := int(math.Floor(i + j))

				if s[prev] == s[next] {
					str := s[prev : next+1]
					if len(childString) < len(str) {
						childString = str
					}
				}
			}
		}
	}

	return childString
}

// @lc code=end
