package leetcode

// 最长的重复子数组
//
// 动态规划
//
// 推理过程:
// 1) dp[i][j] 表示 A[i:] 和 B[j:] 的最长公共前缀;
// 2) dp[i][j] = 0, A[i] != B[j]
// 3) dp[i][j] = dp[i+1][j+1] + 1, A[i] == B[j]
//
// 辅助数组:
// 1) “” 和 A, “” 和 B 公共前缀长度为 0;
//

func findLength(nums1 []int, nums2 []int) int {

	var dp [][]int
	for i := 0; i < len(nums1)+1; i++ {
		dp = append(dp, make([]int, len(nums2)+1))
	}

	var ans int
	for i := len(nums1) - 1; i >= 0; i-- {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}

			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}

	return ans
}
