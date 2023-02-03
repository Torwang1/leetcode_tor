package leetcode

// 题目: 最长严格递增子序列的长度.
//
// 解法: 动态规划;
//
// 递推方程:
// - dp(i) 表示以 nums[i] 结尾的最长严格递增子序列长度;
// - dp(i) = max(dp(j)+1), j < i && nums[j] < nums[i];
// - dp(0) = 1;
//

func lengthOfLIS(nums []int) int {

	dp := make([]int, len(nums))
	dp[0] = 1

	ans := 1
	for i := 1; i < len(dp); i++ {
		var max int
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j] > max {
				max = dp[j]
			}
		}

		dp[i] = max + 1
		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return ans
}
