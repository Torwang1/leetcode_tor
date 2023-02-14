package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := [][]int{}

	current := []*TreeNode{root}

	for i := 0; len(current) != 0; i++ {
		var next []*TreeNode

		for _, node := range current {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}

		var record []int
		if i%2 == 0 {
			for i := 0; i < len(current); i++ {
				record = append(record, current[i].Val)
			}
		} else {
			for i := len(current) - 1; i >= 0; i-- {
				record = append(record, current[i].Val)
			}
		}
		ans = append(ans, record)

		current = next
	}

	return ans
}
