/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */
package main

import (
	"math"
	"sort"
)

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums []int = append(nums1, nums2...)
	sort.Ints(nums)
	var mid float64 = float64(len(nums)) / 2
	if float64(int(mid)) < mid {
		return float64(nums[int(mid)])
	}

	return float64(nums[int(math.Ceil(mid))]+nums[int(math.Ceil(mid-1))]) / 2
}

// @lc code=end
