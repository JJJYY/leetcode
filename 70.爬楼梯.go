/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		if i <= 2 {
			dp[i] = i
			continue
		}

		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n]
}

// @lc code=end

