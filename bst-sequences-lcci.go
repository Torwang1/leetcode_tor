package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BSTSequences(root *TreeNode) [][]int {
	ans := [][]int{}

	if root == nil {
		return ans
	}

	appendAns := func(arr []*TreeNode) {
		one := make([]int, len(arr))
		for i, node := range arr {
			one[i] = node.Val
		}
		ans = append(ans, one)
	}

	arr := make([]*TreeNode, 0)
	visited := make(map[*TreeNode]bool)

	// 初始化 root 节点
	arr = append(arr, root)
	visited[root] = true

	sequence(appendAns, arr, visited)

	return ans
}

// sequence 回溯, 每次填充一个节点.
func sequence(appendAns func(arr []*TreeNode), arr []*TreeNode, visited map[*TreeNode]bool) {
	subnodes := subNodes(arr, visited)

	// 所有节点都被访问过, arr 是一个可用序列;
	if len(subnodes) == 0 {
		appendAns(arr)
		return
	}

	for _, next := range subnodes {
		visited[next] = true
		arr = append(arr, next)
		sequence(appendAns, arr, visited)
		visited[next] = false
		arr = arr[:len(arr)-1] // TODO: 应该是 [...)
	}
}

// subNodes 返回所有 “未被访问过的” 字节点;
func subNodes(arr []*TreeNode, visited map[*TreeNode]bool) []*TreeNode {
	result := make([]*TreeNode, 0, len(arr)*2)

	for _, next := range arr {
		if left := next.Left; left != nil && !visited[left] {
			result = append(result, left)
		}
		if right := next.Right; right != nil && !visited[right] {
			result = append(result, right)
		}
	}

	return result
}
