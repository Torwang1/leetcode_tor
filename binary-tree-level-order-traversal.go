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

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		next := []*TreeNode{}

		for _, node := range queue {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}

		var result []int
		for _, node := range queue {
			result = append(result, node.Val)
		}
		ans = append(ans, result)

		queue = next
	}

	return ans
}
