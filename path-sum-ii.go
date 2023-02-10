package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 题目: 从根节点到叶子节点, 路径总和等于 target 的所有路径.
//
// 基础算法: DFS
//
func pathSum(root *TreeNode, targetSum int) [][]int {
	var ans [][]int

	pathSumii(root, targetSum, 0, &[]int{}, &ans)

	return ans
}

func pathSumii(root *TreeNode, target int, sum int, record *[]int, ans *[][]int) {
	if root == nil {
		return
	}

	sum = sum + root.Val

	*record = append(*record, root.Val)

	if root.Left == nil && root.Right == nil {
		if sum == target {
			cp := make([]int, len(*record))
			copy(cp, *record)
			*ans = append(*ans, cp)
		}
	}

	pathSumii(root.Left, target, sum, record, ans)
	pathSumii(root.Right, target, sum, record, ans)

	*record = (*record)[:len(*record)-1]
}
