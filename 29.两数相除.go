/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */
package divide

// @lc code=start
func divide(dividend int, divisor int) int {
	result := dividend / divisor

	if result >= 2147483648 {
		return result - 1
	}

	return result
}

// @lc code=end
