package leetcode

// 二叉树的层次遍历

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := [][]int{}

	// current: 当前层次
	// next: 下一层
	current := []*TreeNode{root}

	for len(current) > 0 {
		next := []*TreeNode{}

		for _, node := range current {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}

		var result []int
		for _, node := range current {
			result = append(result, node.Val)
		}
		ans = append(ans, result)

		current = next
	}

	return ans
}
