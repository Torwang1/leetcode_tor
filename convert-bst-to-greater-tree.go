package leetcode

// 题目: 把二叉搜索树, 转换为累加树.
//
// 1) 二叉搜索树, 中序有序;
// 2) 前缀和;
//

func convertBST(root *TreeNode) *TreeNode {
	nodes := []*TreeNode{}
	infixOrderTraverse(root, &nodes)

	pre := 0
	for i := len(nodes) - 1; i >= 0; i-- {
		pre += nodes[i].Val
		nodes[i].Val = pre
	}

	return root
}

// 中序遍历二叉树
func infixOrderTraverse(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}

	infixOrderTraverse(root.Left, nodes)
	*nodes = append(*nodes, root)
	infixOrderTraverse(root.Right, nodes)
}
