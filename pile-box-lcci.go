package leetcode

import (
	"math"
	"sort"
)

// pileBox 堆箱子问题.
//
// 思路:
// - 对任意元素做“升序”排序, 最终结果一定是它的“升序子序列”;
//
// dp[i] := max(dp[k]...) + hight(i), 其中 k 满足: k < i && box[k][wi,di,hi] < box[i][wi,di,hi];
func pileBox(box [][]int) int {
	var maxSum int = math.MinInt

	dp := make([]int, len(box))

	maxInt := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dpFunc := func(i int) int {
		var m int = 0
		for k := 0; k < i; k++ {
			if box[k][0] < box[i][0] && box[k][1] < box[i][1] && box[k][2] < box[i][2] {
				m = maxInt(m, dp[k])
			}
		}
		return m + box[i][2]
	}

	sort.Slice(box, func(i, j int) bool { return box[i][0] < box[j][0] })

	for i := 0; i < len(box); i++ {
		dp[i] = dpFunc(i)
		maxSum = maxInt(maxSum, dp[i])
	}

	return maxSum
}
