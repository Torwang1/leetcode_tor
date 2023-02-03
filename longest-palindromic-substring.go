package leetcode

import "fmt"

// 题目: 最长回文子串
//
// 动态规划
//
// 递推方程:
// - dp(i, j) = dp(i+1, j-1) && s[i] == s[j]
// - dp(i, i) = true
// - dp(i, i+1) = true, s[i] == s[i+1]
//
// max := j + 1 - i;
//

func longestPalindrome(s string) string {

	n := len(s)

	dp := [][]bool{}
	for i := 0; i < n; i++ {
		dp = append(dp, make([]bool, n))
	}

	// - dp(i, i) = true
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	// - dp(i, i+1) = true, s[i] == s[i+1]
	for i := 0; i < n; i++ {
		if i+1 < n && s[i] == s[i+1] {
			dp[i][i+1] = true
		}
	}

	// - dp(i, j) = dp(i+1, j-1) && s[i] == s[j]
	//
	// dp(i, j) 依赖“左下角”元素. 所以, 求解顺序从下向上
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			if dp[i+1][j-1] && s[i] == s[j] {
				dp[i][j] = true
			}
		}
	}

	// 遍历 dp 数组, 找出 max := j + 1 - i;
	//
	// 注意: j >= i 的元素有效
	max := 0
	mi, mj := 0, 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			cur := j + 1 - i
			if dp[i][j] && cur > max {
				max = cur
				mi, mj = i, j
			}
		}
	}
	return s[mi : mj+1]
}

func printTwoDimensionalArray(array [][]bool) {
	for _, line := range array {
		fmt.Println(line)
	}
}
