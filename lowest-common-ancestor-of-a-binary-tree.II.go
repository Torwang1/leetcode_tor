package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 题目: 二叉树的最近公共祖先
//
// 官方题解:
// 1) f(x) 表示 x 的子树中是否包含 p 或 q 节点;
//
// 最近公共祖先 := ( f(x.left) && f(x.right) ) || ( (x==q || x==p) && (f(x.left) || f(x.right)) ), 其中:
// 1) f(x.left) && f(x.right): x 的左右子树中分别包含 p, q, 则 x 是公共祖先;
// 2) (x==q || x==p) && (f(x.left) || f(x.right)): x 等于 p 或 q, x 的子树中包含 p 或 q, 则 x 是公共祖先;
//
// 官方题解的推理, 无法从公式直接映射到代码.
//
// 直观理解:
// 1) p, q 一定在树中;
// 2) 对于 root = p 或 q 的场景, 直接返回 root 节点;
// 3) 对于 root != p 或 q 的场景, 如果 root.left root.right 分别包含 p 和 q, 则返回 root.
//

func lowestCommonAncestorII(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestorII(root.Left, p, q)
	right := lowestCommonAncestorII(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}
