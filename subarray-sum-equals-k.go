package leetcode

// 题目: 和为 k 的子数组
//
// 核心算法: 前缀和 + hash
//
// 递推公式:
// 1) sum[i,j] = pre(j) - pre(i-1) = k, i <= j
// 2) pre(i-1) = pre(j) - k, 以 j 结尾, 有多少 (i-1) 使得等式成立

func subarraySum(nums []int, k int) int {
	var ans int

	pre := 0                  // 前缀和
	hash := map[int]int{0: 1} // 记录 pre = xxx 的个数

	for j := 0; j < len(nums); j++ {
		pre += nums[j]

		target := pre - k
		ans += hash[target]

		hash[pre] += 1
	}

	return ans
}
