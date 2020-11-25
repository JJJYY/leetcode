/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */
package subsets

// @lc code=start
func subsets(nums []int) [][]int {
	result := make([][]int)

	for i := 0; i < len(nums); i++ {
		child := make([]int)
		child = append(child, nums[i])

		for j := 0; j <= len(nums); j++ {
			child = append(child, nums[j])
		}
	}

	return result
}

// @lc code=end
