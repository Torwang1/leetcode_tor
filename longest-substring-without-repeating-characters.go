package leetcode

// 题目: 无重复字符的最长子串.
//
// 滑动窗口 /或者说是/ 双指针.

func lengthOfLongestSubstring(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var ans int

	// 计数器的数值, 只可能是 0 或 1.
	counter := map[byte]int{}

	// 左闭右开区间, 方便计算元素数量.
	for l, r := 0, 0; r < len(s); /**/ {
		c := s[r]

		for counter[c] > 0 {
			counter[s[l]] = 0 // 左边界右移.
			l++
		}

		counter[c] = 1 // 右边界右移.
		r++

		ans = max(ans, r-l)
	}

	return ans
}
