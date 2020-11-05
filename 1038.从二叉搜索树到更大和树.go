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

func getLeftNode(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}

	return root
}

func bstToGst(root *TreeNode, base ...int) *TreeNode {
	val := 0
	if len(base) > 0 {
		val = base[0]
	}

	if root.Right != nil {
		right := bstToGst(root.Right, val)

		if right.Left != nil {
			left := getLeftNode(right)
			root.Val = left.Val + root.Val
		} else {
			root.Val = right.Val + root.Val
		}
	} else {
		root.Val = root.Val + val
	}

	if root.Left != nil {
		bstToGst(root.Left, root.Val)
	}

	return root
}

// @lc code=end
