/*
 * @lc app=leetcode.cn id=1038 lang=golang
 *
 * [1038] 从二叉搜索树到更大和树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func bstToGst(root *TreeNode, base ...int) *TreeNode {
	val := 0
	if len(base) > 0 {
		val = base[0]
	}

	if root.Right != nil {
		right := bstToGst(root.Right, val)

		if right.Left != nil {
			root.Val = right.Left.Val + root.Val + val
		} else {
			root.Val = right.Val + root.Val
		}
	}

	if root.Left != nil {
		root.Left.Val = root.Val + root.Left.Val
		bstToGst(root.Left, root.Val+val)
	}

	root.Val = root.Val + val

	return root
}

// @lc code=end
