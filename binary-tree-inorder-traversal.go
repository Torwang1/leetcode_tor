package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 基础算法: 二叉树中序遍历

func inorderTraversal(root *TreeNode) []int {

	var inorder func(root *TreeNode, ans *[]int)
	inorder = func(root *TreeNode, ans *[]int) {
		if root == nil {
			return
		}

		inorder(root.Left, ans)
		*ans = append(*ans, root.Val)
		inorder(root.Right, ans)
	}

	var ans []int

	inorder(root, &ans)

	return ans
}
