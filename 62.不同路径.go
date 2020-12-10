/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	if m <= 1 || n <= 1 {
		return 1
	}

	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
}

// @lc code=end

