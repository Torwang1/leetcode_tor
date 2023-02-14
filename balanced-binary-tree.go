package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 判断平衡树

func isBalanced(root *TreeNode) bool {
	var dfs func(root *TreeNode) (int, bool)

	abs := func(a int) int {
		if a > 0 {
			return a
		}
		return -1 * a
	}

	dfs = func(root *TreeNode) (int, bool) {
		if root == nil {
			return 0, true
		}

		ldepth, lbool := dfs(root.Left)
		rdepth, rbool := dfs(root.Right)

		depth := ldepth
		if depth < rdepth {
			depth = rdepth
		}
		depth++

		if lbool == false || rbool == false || abs(ldepth-rdepth) > 1 {
			return depth, false
		}
		return depth, true
	}

	_, b := dfs(root)

	return b
}
