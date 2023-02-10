package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 题目: “从 根 到 叶子” 所有路径上, 子路径和为 target 的数量;
//
// 基础算法: DFS + 前缀和
//

func pathSumiii(root *TreeNode, targetSum int) int {
	var ans int

	pathSumiiiDFS(root, targetSum, &ans, 0, map[int]int{0: 1})

	return ans
}

// pathSumiiiDFS 求解子路径和为 target 的数量.
//
// 输入参数:
// 1) ans 记录 “以 root 结尾” 并 “满足子路径和等于 target” 的数量;
// 2) pre 前缀和;
// 3) hash 前缀和 = k 的个数. 其中, 不包含 root 结尾的前缀和;
func pathSumiiiDFS(root *TreeNode, target int, ans *int, pre int, hash map[int]int) {
	if root == nil {
		return
	}

	pre = pre + root.Val

	*ans += hash[pre-target]

	hash[pre] += 1

	pathSumiiiDFS(root.Left, target, ans, pre, hash)
	pathSumiiiDFS(root.Right, target, ans, pre, hash)

	hash[pre] -= 1
}
