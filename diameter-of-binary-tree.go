package leetcode

// 题目: 二叉树的直径, 任意两节点的路径长度, 最大值.
//
// 类似: 平衡树, 深度
//
// 直径 := 左深度 + 右深度
// 深度 := max(左深度, 右深度) + 1
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	var ans int

	maxDiameter(root, &ans)

	return ans
}

// maxDiameter 求解二叉树直径.
//
// 输入参数:
// 1) ans 记录递归过程中的最大值;
//
// 返回值:
// 1) depth 表示: 当前节点深度;
func maxDiameter(root *TreeNode, ans *int) (depth int) {
	if root == nil {
		return 0
	}

	left := maxDiameter(root.Left, ans)
	right := maxDiameter(root.Right, ans)

	diameter := left + right
	if diameter > *ans {
		*ans = diameter
	}

	if left > right {
		return left + 1
	}
	return right + 1
}
