package leetcode

// 问题: 回文子串的数量;
//
// 与问题“最长回文子串”相同, 使用动态规划求解;
//
// 递推方程:
// 1) dp[i][j] = dp[i+1][j-1] ^ s[i] == s[j]
// 2) dp[i][i] = true
// 3) dp[i][i+1] = true, s[i] == s[i+1]
//
// 第一反应:
// 1) 子串, 立刻想到“下标”[i,j]

func countSubstrings(s string) int {
	var ans int

	var dp [][]bool
	for i := 0; i < len(s); i++ {
		dp = append(dp, make([]bool, len(s)))
	}

	// 2) dp[i][i] = true
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
		ans++
	}

	// 3) dp[i][i+1] = true, s[i] == s[i+1]
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			ans++
		}
	}

	// 1) dp[i][j] = dp[i+1][j-1] ^ s[i] == s[j]
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 2; j < len(s); j++ {
			if dp[i+1][j-1] && s[i] == s[j] {
				dp[i][j] = true
				ans++
			}
		}
	}

	return ans
}
