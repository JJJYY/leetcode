/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */
package main

import (
	"strconv",
	"strings"
)

// @lc code=start
func numDecodings(s string) int {
	s = strings.Trim(s, "0")
	
	length := len(s)
	if length <= 1 {
		l, _ := strconv.Atoi(s)
		if l <= 0 {
			return 0
		}

		return length
	}

	num := 0
	i, _ := strconv.Atoi(s[length-2 : length])
	if i <= 26 && i > 0 {
		num = 1
	}

	return numDecodings(s[0:length-1]) + num
}

// func main() {
// 	numDecodings("12")
// }

// @lc code=end
