package leetcode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	var isValid func(root *TreeNode, lower int, upper int) bool

	isValid = func(root *TreeNode, lower, upper int) bool {
		if root == nil {
			return true
		}
		if root.Val <= lower || root.Val >= upper {
			return false
		}
		return isValid(root.Left, lower, root.Val) && isValid(root.Right, root.Val, upper)
	}

	return isValid(root, math.MinInt, math.MaxInt)
}
