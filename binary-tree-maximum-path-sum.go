package leetcode

import "math"

// 二叉树中的最大路径和
//
// 路径: 任意节点到另一节点.
//
// maxGain(root): 以 root 为起点, 子树中的最大路径和;
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var maxGain func(root *TreeNode) int
	maxGain = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		// 子树的最大贡献值, 大于等于 0.
		left := max(maxGain(root.Left), 0)
		right := max(maxGain(root.Right), 0)

		// 记录最大路径和.
		ans = max(ans, root.Val+left+right)

		// 以 root 为起点的最大贡献值.
		return root.Val + max(left, right)
	}

	maxGain(root)

	return ans
}
