/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
package threeSum

import (
	"fmt"
	"sort"
)

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	cache := map[string]int{}
	result := make([][]int, 0)
	length := len(nums)
	for i1 := 0; i1 < length-2; i1++ {
		for i2 := i1 + 1; i2 < length-1; i2++ {
			i3Val := (nums[i1] + nums[i2]) * -1
			for i3 := i2 + 1; i3 < length; i3++ {
				key := fmt.Sprint("[%v%v%v]", nums[i1], nums[i2], nums[i3])

				if i3Val == nums[i3] && cache[key] < 1 {
					result = append(result, []int{nums[i1], nums[i2], nums[i3]})

					cache[key] = 1
				}
			}
		}
	}

	return result
}

// func main() {
// 	test := []int{34, 55, 79, 28, 46, 33, 2, 48, 31, -3, 84, 71, 52, -3, 93, 15, 21, -43, 57, -6, 86, 56, 94, 74, 83, -14, 28, -66, 46, -49, 62, -11, 43, 65, 77, 12, 47, 61, 26, 1, 13, 29, 55, -82, 76, 26, 15, -29, 36, -29, 10, -70, 69, 17, 49}

// 	threeSum(test)
// }

// @lc code=end
