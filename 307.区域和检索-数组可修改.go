/*
 * @lc app=leetcode.cn id=307 lang=golang
 *
 * [307] 区域和检索 - 数组可修改
 */
package search

// @lc code=start
type NumArray struct {
	Data []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		Data: nums,
	}
}

func (that *NumArray) Update(i int, val int) {
	that.Data[i] = val
}

func (that *NumArray) SumRange(i int, j int) int {
	sum := 0

	for ; i <= j; i++ {
		sum += that.Data[i]
	}

	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
// @lc code=end
