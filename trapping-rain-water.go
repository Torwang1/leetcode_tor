package leetcode

// 题目: 接雨水
//
// 题目分析:
// 1) i 能接多少雨水, 取决于: leftMax 和 rightMax;
//
// 朴素解法:
// 1) 在下标 i, 分别向左和向右扫描, 获得 leftMax 和 rightMax. trap := max( min(leftMax, rightMax) - height[i], 0 )
//
// 动态规划:
// 1) 预处理, 获取 leftMost 和 rightMost 数组;
// 2) 其中, leftMost[i] 表示: 从左向右, 下标i及i的左侧, 最高的柱子;
//

func trap(height []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := len(height)

	leftMost := make([]int, n)
	leftMost[0] = height[0]
	for i := 1; i < n; i++ {
		leftMost[i] = max(leftMost[i-1], height[i])
	}

	rightMost := make([]int, n)
	rightMost[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMost[i] = max(rightMost[i+1], height[i])
	}

	var sum int
	for i := 0; i < n; i++ {
		sum += max(min(leftMost[i], rightMost[i])-height[i], 0)
	}

	return sum
}
