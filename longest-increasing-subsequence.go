package leetcode

import "math"

func lengthOfLIS(nums []int) int {
	var max int = math.MinInt

	dp := make([]int, len(nums))

	maxInt := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dpFunc := func(i int) int {
		var m int = 0
		for k := 0; k < i; k++ {
			if nums[k] < nums[i] {
				m = maxInt(m, dp[k])
			}
		}
		return m + 1
	}

	for i := 0; i < len(nums); i++ {
		dp[i] = dpFunc(i)
		max = maxInt(max, dp[i])
	}

	return max
}
