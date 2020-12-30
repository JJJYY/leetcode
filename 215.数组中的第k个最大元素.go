/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */
package findKthLargest

import (
	"sort"
)

// @lc code=start
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })

	max := nums[0]
	for _, item := range nums {
		if (item == max) {
		    k -= 1
			continue
		}

		if (k == 1) {
            max = item
			return item
		}
        
        k -= 1
	}

    return nums[0]
}
// @lc code=end

