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

	reverse := false
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		var next []*TreeNode

		for _, node := range queue {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}

		var record []int
		if reverse {
			for i := len(queue) - 1; i >= 0; i-- {
				record = append(record, queue[i].Val)
			}
		} else {
			for i := 0; i < len(queue); i++ {
				record = append(record, queue[i].Val)
			}
		}
		ans = append(ans, record)

		queue = next
		switch reverse {
		case true:
			reverse = false
		case false:
			reverse = true
		}
	}

	return ans
}
