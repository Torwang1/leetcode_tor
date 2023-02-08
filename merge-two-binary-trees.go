package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 题目: 合并二叉树
//
// DFS, 栈
//
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	node := &TreeNode{}

	switch {
	case root1 == nil && root2 == nil:
		return nil
	case root1 == nil && root2 != nil:
		node.Val = root2.Val
	case root1 != nil && root2 == nil:
		node.Val = root1.Val
	default:
		node.Val = root1.Val + root2.Val
	}

	var l1, l2 *TreeNode
	if root1 != nil {
		l1 = root1.Left
	}
	if root2 != nil {
		l2 = root2.Left
	}
	node.Left = mergeTrees(l1, l2)

	var r1, r2 *TreeNode
	if root1 != nil {
		r1 = root1.Right
	}
	if root2 != nil {
		r2 = root2.Right
	}
	node.Right = mergeTrees(r1, r2)

	return node
}
