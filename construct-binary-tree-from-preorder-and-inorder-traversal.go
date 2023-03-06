package leetcode

// 从前序遍历和中序遍历, 构造二叉树.
//
// 前序: 根, 左子树, 右子树;
// 中序: 左子树, 根, 右子树;
//
// 前序: 确定根元素;
// 中序: 确定“左/右子树”长度;
//
// 优化点:
// 1) 使用 hash 简化根元素查找的复杂度;
//
// 基础算法:
// 1) 二叉树的递归特性;
//
// 题目约束条件:
// 1) 元素各不相同;
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	hash := make(map[int]int, len(inorder))
	for i, v := range inorder {
		hash[v] = i
	}

	value := preorder[0]
	index := hash[value]

	// index: 0, 1, 2, 3, 4
	//              |        index = 2
	//
	// index 表示: 根在中序数组中的下标;
	// index 表示: 左子树的长度;
	// index+1 表示: 右子树的起点;

	root := &TreeNode{
		Val: value,
		Left: buildTree(
			preorder[1:index+1],
			inorder[0:index],
		),
		Right: buildTree(
			preorder[index+1:],
			inorder[index+1:],
		),
	}

	return root
}
