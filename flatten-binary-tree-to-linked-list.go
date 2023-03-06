package leetcode

// 二叉树展开为“前序链表”
//
// 处理过程:
// 1) 摘除 root.right = tmp;
// 2) root.right = root.left, root.left = nil;
// 3) 寻找 root 的 rightmost, rightmost.right = tmp;
//
// 4) root = root.right, root == nil 跳出循环.
//
// 处理原则:
// 1) 任何一次变换后, 得到的前序序列相同;

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	for root != nil {
		node := root.Right

		root.Right = root.Left
		root.Left = nil

		rightmost := root
		for rightmost.Right != nil {
			rightmost = rightmost.Right
		}
		rightmost.Right = node

		root = root.Right
	}
}
