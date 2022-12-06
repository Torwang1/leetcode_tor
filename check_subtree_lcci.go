package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// checkSubTree 暴力解法: 遍历 t1, 逐个节点与 t2 比较子树;
func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	return subTree(t1, t2)
}

// subTree 前序深度便利.
func subTree(t1, t2 *TreeNode) bool {
	if ok := matchTree(t1, t2); ok {
		return true
	}

	if t1 != nil {
		return subTree(t1.Left, t2) || subTree(t1.Right, t2)
	}

	return false
}

// matchTree 递归匹配
func matchTree(t1 *TreeNode, t2 *TreeNode) bool {
	switch {
	case t1 == nil && t2 == nil:
		return true
	case t1 != nil && t2 != nil:
		return t1.Val == t2.Val && matchTree(t1.Left, t2.Left) && matchTree(t1.Right, t2.Right)
	default:
		return false
	}
}
