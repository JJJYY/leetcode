/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */
package maxSubArray

// @lc code=start
func maxSubArray(nums []int) int {
	sum := nums[0]

	for i := 0; i < len(nums); i++ {
		s := 0
		for j := i; j < len(nums); j++ {
			s += nums[j]

			if s > sum {
				sum = s
			}
		}
	}

	return sum
}

// @lc code=end
