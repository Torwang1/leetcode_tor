package leetcode

// 从后序遍历和中序遍历, 构造二叉树.
//
// 后序: 左子树, 右子树, 根;
// 中序: 左子树, 根, 右子树;
//
// 后序: 确定根元素;
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
func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	hash := make(map[int]int, len(inorder))
	for i, v := range inorder {
		hash[v] = i
	}

	value := postorder[len(postorder)-1]
	index := hash[value]

	// index: 0, 1, 2, 3, 4
	//              |        index = 2
	//
	// index 表示: 根在中序数组中的下标;
	// index 表示: 左子树的长度;
	// 左子树的起始点: 相同;

	root := &TreeNode{
		Val: value,
		Left: buildTree(
			inorder[0:index],
			postorder[0:index],
		),
		Right: buildTree(
			inorder[index+1:],
			postorder[index:len(postorder)-1],
		),
	}

	return root
}
