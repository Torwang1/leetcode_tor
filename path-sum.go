package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 题目: 从根节点到叶子节点, 路径总和等于 target.
//
// 基础算法: DFS
//
func hasPathSum(root *TreeNode, targetSum int) bool {
	return hasPathSumDFS(root, targetSum, 0)
}

func hasPathSumDFS(root *TreeNode, target int, sum int) bool {
	if root == nil {
		return false
	}

	sum = sum + root.Val

	if root.Left == nil && root.Right == nil {
		if sum == target {
			return true
		}
	}

	if ok := hasPathSumDFS(root.Left, target, sum); ok {
		return true
	}
	if ok := hasPathSumDFS(root.Right, target, sum); ok {
		return true
	}

	return false
}
