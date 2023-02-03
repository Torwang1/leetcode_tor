package leetcode

// 题目: 二叉树, 两个节点的共同祖先
//
// 基础算法: 后序遍历 + visited

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode

	traverse(root, p, q, &ans)

	return ans
}

// 后序遍历(非完整实现)
func traverse(root, p, q *TreeNode, ans **TreeNode) (pf, qf bool) {

	var lpf, lqf bool
	if root.Left != nil {
		lpf, lqf = traverse(root.Left, p, q, ans)
		if *ans != nil {
			return
		}
	}

	var rpf, rqf bool
	if root.Right != nil {
		rpf, rqf = traverse(root.Right, p, q, ans)
		if *ans != nil {
			return
		}
	}

	pf = lpf || rpf || root == p
	qf = lqf || rqf || root == q

	if pf && qf {
		*ans = root
	}
	return pf, qf
}
