/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
 */
package invertTree

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	Left := root.Left
	Right := root.Right

	if Right != nil {
		root.Left = invertTree(Right)
	} else {
		root.Left = nil
	}

	if Left != nil {
		root.Right = invertTree(Left)
	} else {
		root.Right = nil
	}

	return root
}

// @lc code=end
