package leetcode

// 题目: 二叉树的右视图;
//
// 解法: BFS 最左元素;

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ans := []int{}

	nodes := []*TreeNode{root}

	for len(nodes) != 0 {
		ans = append(ans, nodes[len(nodes)-1].Val)

		next := []*TreeNode{}
		for _, node := range nodes {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		nodes = next
	}

	return ans
}
