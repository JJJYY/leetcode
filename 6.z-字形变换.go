/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */
package main

import "strings"

// @lc code=start
func convert(str string, numRows int) string {
	if numRows == 1 {
		return str
	}

	strArr := make([]string, numRows)

	row := 0
	dir := 1
	for i := range str {
		strArr[row] = strArr[row] + string(str[i])

		if row == numRows-1 {
			dir = -1
		} else if row == 0 {
			dir = 1
		}

		row += dir
	}

	return strings.Join(strArr, "")
}

// @lc code=end
