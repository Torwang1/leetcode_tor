package leetcode

// 题目: 从两个数组中, 挑选子序列, 组合成最大数字;
//
// 分治 + 单一数组最大子序列(单调栈);
//
// 记录分治最大值;
//
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	ans := make([]int, k)

	for i := 0; i <= k; i++ {
		f1, f2 := i, k-i
		if f1 > len(nums1) || f2 > len(nums2) {
			continue
		}

		merged := merge(maxSubsequence(nums1, f1), maxSubsequence(nums2, f2))
		if lexicographicalLess(ans, merged) {
			ans = merged
		}
	}

	return ans
}

// maxSubsequence 最大子序列, 单调栈实现;
//
// 注意: 这里有两种写法: “保留 k 个” 和 “删除 k 个”;
func maxSubsequence(nums []int, k int) []int {
	ans := make([]int, 0, k)

	for i, v := range nums {
		for len(ans) > 0 && v > ans[len(ans)-1] && len(nums)-(i+1)+len(ans) >= k {
			ans = ans[:len(ans)-1]
		}
		ans = append(ans, v)
	}

	// 边界条件: nums 递减序列, 会导致 len(ans) > k 的情况;
	return ans[:k]
}

// merge 大数尽量往前放.
func merge(a, b []int) []int {
	ans := make([]int, len(a)+len(b))

	for i := 0; i < len(ans); i++ {
		if lexicographicalLess(a, b) {
			ans[i], b = b[0], b[1:]
		} else {
			ans[i], a = a[0], a[1:]
		}
	}

	return ans
}

// lexicographicalLess 字典序比较. 兼容: 长度不一致的情况.
func lexicographicalLess(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}
