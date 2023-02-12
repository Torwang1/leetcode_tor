package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 题目: 二叉搜索树中, 第 K 个最小元素.
//
// 中序有序;
//
// 基础算法: DFS
//
// left;
// root;
// right;

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		k--

		if k == 0 {
			return root.Val
		}

		root = root.Right
	}

}
