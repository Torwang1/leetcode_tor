package leetcode

// 问题: 目标和
//
// 数组中的任意元素组合, sum = target;
//
// 动态规划. 递推方程:
// 1) dp[i][j] 前 i 个元素任意组合, 结果为 j 的方案数量;
// 2) dp[i][j] = dp[i-1][j], j < nums[i]
// 3) dp[i][j] = dp[i-1][j] + dp[i-1][j-num[i]], j >= nums[i]
// 4) dp[0][0] = 1
// 5) dp[0][j] = 0, j > 0
//

func findTargetSumWays(nums []int, target int) int {

	// 转换过程
	var sum int
	for _, v := range nums {
		sum += v
	}

	// 求解数组中任意元素之和, 等于 neg.
	diff := sum - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	neg := (sum - target) / 2

	n := len(nums)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, neg+1)
	}

	// 4) dp[0][0] = 1
	// 5) dp[0][j] = 0, j > 0
	dp[0][0] = 1
	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = 0
	}

	// 2) dp[i][j] = dp[i-1][j], j < nums[i]
	// 3) dp[i][j] = dp[i-1][j] + dp[i-1][j-num[i]], j >= nums[i]
	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if j < nums[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[n][neg]
}
