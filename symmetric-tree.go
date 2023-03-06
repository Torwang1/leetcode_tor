package leetcode

// 对称二叉树. 根节点为轴, 左右对称

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricImpl(root.Right, root.Left)
}

func isSymmetricImpl(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if (a == nil && b != nil) ||
		(a != nil && b == nil) ||
		a.Val != b.Val {
		return false
	}

	if ok := isSymmetricImpl(a.Right, b.Left); !ok {
		return false
	}
	return isSymmetricImpl(a.Left, b.Right)
}
