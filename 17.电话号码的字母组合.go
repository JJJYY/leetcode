/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */
package main

import (
	"strconv"
)

// @lc code=start
func letterCombinations(digits string) []string {
	numberArray := [9]string{
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}

	var result = make([]string, 0)

	for _, s := range digits {
		num, _ := strconv.Atoi(string(s))
		var numStr = numberArray[num-1]
		var newResult = make([]string, 0)

		if len(result) > 0 {
			for _, s2 := range result {
				for _, s1 := range numStr {
					newResult = append(newResult, s2+string(s1))
				}
			}
		} else {
			for _, s1 := range numStr {
				newResult = append(newResult, string(s1))
			}
		}

		result = newResult
	}

	return result
}

// @lc code=end
